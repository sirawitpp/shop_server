package mock

import (
	"errors"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/errs"
)

type UserRepositoryMock struct{}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) Register(input model.User) (*model.User, error) {
	switch input.Username {
	case "pass":
		return &input, nil
	case "username":
		return nil, errors.New(errs.UsernameAlreadyExists + " " + errs.SQLSTATE23505)
	case "email":
		return nil, errors.New(errs.EmailAlreadyExists + " " + errs.SQLSTATE23505)
	default:
		return nil, errors.New("")
	}

}
