package repository

import (
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		input := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomUsername(),
			Email:    random.RandomEmail(),
		}
		user, err := testUserQuery.Register(input)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user.Username, input.Username)
		assert.Equal(t, user.Email, input.Email)
		assert.NotNil(t, user.CreatedAt)
		assert.Greater(t, user.ID, uint64(0))
	})
	t.Run("failed (user alreaydy exists)", func(t *testing.T) {
		input := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomUsername(),
			Email:    random.RandomEmail(),
		}
		_, _ = testUserQuery.Register(input)
		user2, err := testUserQuery.Register(input)
		assert.Empty(t, user2)
		assert.Error(t, err)
	})

	t.Run("failed (email alreaydy exists)", func(t *testing.T) {
		input := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomUsername(),
			Email:    random.RandomEmail(),
		}
		_, _ = testUserQuery.Register(input)
		input2 := input
		input2.Username = random.RandomUsername()
		user2, err := testUserQuery.Register(input2)
		assert.Empty(t, user2)
		assert.Error(t, err)
	})
}
