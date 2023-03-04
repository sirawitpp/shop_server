package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/converter"
	"sirawit/shop/pkg/pb"
)

func (u *UserServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	user, err := u.userService.Register(model.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	token, _ := u.tokenManager.CreateToken(user.Username, u.config.TokenDuration)
	return converter.ConvertUserToRegisterRes(user, token), nil
}
