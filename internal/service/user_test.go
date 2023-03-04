package service

import (
	"errors"
	"fmt"
	"sirawit/shop/internal/model"
	"sirawit/shop/mock"
	"sirawit/shop/pkg/errs"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		a := mock.NewUserRepositoryMock()
		b := NewUserService(a)
		user := model.User{
			ID:       random.RandomUInt64(0, 10),
			Username: random.RandomUsername(),
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}
		a.On("Register", user).Return(&user, nil)
		result, err := b.Register(user)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.Username, user.Username)
		assert.Equal(t, result.Email, user.Email)
		assert.NotNil(t, result.ID)
	})

	t.Run("failed (server error)", func(t *testing.T) {
		a := mock.NewUserRepositoryMock()
		b := NewUserService(a)
		user := model.User{}
		a.On("Register", user).Return(&user, errors.New(""))
		result, err := b.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Internal")
	})

	t.Run("failed (username already exist)", func(t *testing.T) {
		a := mock.NewUserRepositoryMock()
		b := NewUserService(a)
		user := model.User{}
		// testUserService = NewUserService(userRepositoryMock)
		a.On("Register", user).Return(&user, fmt.Errorf("username %v", errs.SQLSTATE23505))
		result, err := b.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), errs.UsernameAlreadyExists)
	})

	t.Run("failed (email already exist)", func(t *testing.T) {
		a := mock.NewUserRepositoryMock()
		b := NewUserService(a)
		user := model.User{}
		// testUserService = NewUserService(userRepositoryMock)
		a.On("Register", user).Return(&user, errors.New(errs.SQLSTATE23505))
		result, err := b.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), errs.EmailAlreadyExists)
	})
}
