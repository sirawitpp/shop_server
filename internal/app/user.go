package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"

	"github.com/rs/zerolog/log"
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

func (u *userServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	result, err := u.userService.Register(model.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	_, err = u.loggerClient.SendLoginTimestampToLogger(ctx, &pb.LoginTimestamp{
		Username:       result.User.Username,
		LoginTimestamp: timestamppb.Now(),
	})
	if err == nil {
		log.Info().Msg("Send to logger service success")
	} else {
		log.Err(err).Msg("Send to logger service fail")
	}
	return convertUserToUserRes(result), nil
}

func (u *userServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	result, err := u.userService.Login(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	_, err = u.loggerClient.SendLoginTimestampToLogger(ctx, &pb.LoginTimestamp{
		Username:       result.User.Username,
		LoginTimestamp: timestamppb.Now(),
	})

	if err == nil {
		log.Info().Msg("Send to logger service success")
	} else {
		log.Err(err).Msg("Send to logger service fail")
	}
	return (*pb.LoginRes)(convertUserToUserRes(result)), nil
}
