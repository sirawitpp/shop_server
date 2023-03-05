package service

import (
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/errs"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRes struct {
	User  model.User
	Token string
}

func (u *userService) Register(input model.User) (*UserRes, error) {
	var err error
	input.Password, err = HashPassword(input.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password %v", err)
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
		return nil, status.Errorf(codes.Internal, "failed to create token %v", err)
	}

	return &UserRes{
		User:  *user,
		Token: token,
	}, nil
}
