package service

import (
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/repository"
	"sirawit/shop/pkg/errs"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Register(input model.User) (*model.User, error)
}

type userService struct {
	db repository.UserQuery
}

func NewUserService(db repository.UserQuery) UserService {
	return &userService{db: db}
}

func (u *userService) Register(input model.User) (*model.User, error) {
	user, err := u.db.Register(input)
	if err != nil {
		if strings.Contains(err.Error(), errs.SQLSTATE23505) {
			if strings.Contains(err.Error(), "username") {
				return nil, status.Error(codes.AlreadyExists, errs.UsernameAlreadyExists)
			}
			return nil, status.Error(codes.AlreadyExists, errs.EmailAlreadyExists)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %v", err)

	}

	return user, err
}
