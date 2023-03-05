package app

import (
	"context"
	"log"
	"sirawit/shop/pkg/pb"
)

type LoggerServer struct {
	pb.UnimplementedLoggerServiceServer
}

func NewLoggerServer() *LoggerServer {
	return &LoggerServer{}
}

func (l *LoggerServer) SendLoginTimestampToLogger(ctx context.Context, req *pb.LoginTimestamp) (*pb.LoginTimestamp, error) {
	log.Print(req.LoginTimestamp.AsTime())
	return req, nil
}
