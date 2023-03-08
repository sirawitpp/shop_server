package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/pb"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (l *loggerServer) SendLoginTimestampToLogger(ctx context.Context, req *pb.LoginTimestamp) (*pb.LoginTimestamp, error) {
	result, err := l.service.InsertLoginTimestamp(model.Logger{
		Username: req.GetUsername(),
	})
	if err != nil {
		log.Err(err).Msg("cannot insert to logger db")
		return nil, err
	}
	log.Info().Msg("insert into logger db success")
	return &pb.LoginTimestamp{
		Username:       result.Username,
		LoginTimestamp: timestamppb.New(result.LoginTimestamp),
	}, nil
}
