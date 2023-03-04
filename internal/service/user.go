package service

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/repository"
	"sirawit/shop/pkg/errs"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Register(input model.User) (*RegisterRes, error)
}

type userService struct {
	db           repository.UserQuery
	tokenManager TokenManager
	config       config.UserConfig
}

func NewUserService(db repository.UserQuery, config config.UserConfig) UserService {
	tokenManager := NewTokenManager(config.Sign)
	return &userService{db: db, config: config, tokenManager: tokenManager}
}

type RegisterRes struct {
	User  model.User
	Token string
}

func (u *userService) Register(input model.User) (*RegisterRes, error) {
	var err error
	input.Password, err = HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
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
	token, err := u.tokenManager.CreateToken(user.Username, u.config.TokenDuration)
	if err != nil {
		return nil, err
	}

	return &RegisterRes{
		User:  *user,
		Token: token,
	}, nil
}
