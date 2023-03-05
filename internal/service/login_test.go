package service

import (
	"errors"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		user := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomString(6),
		}
		hashedPassword, err := HashPassword(user.Password)
		assert.NoError(t, err)
		testDB.On("FindUserByUsername", user.Username).Return(&model.User{Username: user.Username, Password: hashedPassword}, nil)
		result, err := testUserService.Login(user.Username, user.Password)
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.NotEqual(t, user.Username, result.User.Password)
	})
	t.Run("password not match", func(t *testing.T) {
		username := random.RandomUsername()
		password := random.RandomString(6)
		testDB.On("FindUserByUsername", username).Return(&model.User{}, nil)
		result, err := testUserService.Login(username, password)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "incorrect password")
	})

	t.Run("user not found", func(t *testing.T) {
		username := random.RandomUsername()
		password := random.RandomString(6)
		testDB.On("FindUserByUsername", username).Return(&model.User{}, errors.New(""))
		result, err := testUserService.Login(username, password)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}
