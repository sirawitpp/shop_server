package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *userService) Login(username, password string) (*UserRes, error) {
	user, err := u.db.FindUserByUsername(username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if ok := CheckPasswordHash(password, user.Password); !ok {
		return nil, status.Error(codes.NotFound, "incorrect password")
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
