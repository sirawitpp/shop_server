package converter

import (
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUserToRegisterRes(user *model.User, token string) *pb.RegisterRes {
	return &pb.RegisterRes{
		User: &pb.User{
			Username:  user.Username,
			Password:  user.Password,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
		Token: token,
	}
}
