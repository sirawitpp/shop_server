package service

import (
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/errs"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		user := model.User{
			Username: "pass",
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}
		result, err := testUserService.Register(user)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.User.ID)
		assert.NotNil(t, result.Token)
		assert.Equal(t, result.User.Username, user.Username)
		assert.Equal(t, result.User.Email, user.Email)
		assert.NotEqual(t, result.User.Password, user.Password)
	})

	t.Run("failed (server error)", func(t *testing.T) {
		user := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}
		result, err := testUserService.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Internal")
	})

	t.Run("failed (username already exist)", func(t *testing.T) {
		user := model.User{
			Username: "username",
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}
		result, err := testUserService.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), errs.UsernameAlreadyExists)
	})

	t.Run("failed (email already exist)", func(t *testing.T) {
		user := model.User{
			Username: "email",
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}
		result, err := testUserService.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), errs.EmailAlreadyExists)
	})
}
