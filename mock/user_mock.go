package mock

import (
	"sirawit/shop/internal/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) Register(input model.User) (*model.User, error) {
	args := m.Called(input)
	return args.Get(0).(*model.User), args.Error(1)
}
