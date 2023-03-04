package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserToRegisterRes(result *service.RegisterRes) *pb.RegisterRes {
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
	return convertUserToRegisterRes(result), nil
}
