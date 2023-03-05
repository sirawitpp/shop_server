package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserToUserRes(result *service.UserRes) *pb.RegisterRes {
	return &pb.RegisterRes{
		User: &pb.User{
			Username:  result.User.Username,
			Email:     result.User.Email,
			CreatedAt: timestamppb.New(result.User.CreatedAt),
		},
		Token: result.Token,
	}
}

func (u *UserServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	result, err := u.userService.Register(model.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	return convertUserToUserRes(result), nil
}

func (u *UserServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	result, err := u.userService.Login(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return (*pb.LoginRes)(convertUserToUserRes(result)), nil
}
