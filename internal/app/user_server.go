package app

import (
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"

	"google.golang.org/grpc"
)

type userServer struct {
	userService service.UserService
	pb.UnimplementedUserServiceServer
	loggerClient pb.LoggerServiceClient
}

func NewUserServer(userService service.UserService, conn *grpc.ClientConn) *userServer {
	loggerClient := pb.NewLoggerServiceClient(conn)
	return &userServer{
		userService:  userService,
		loggerClient: loggerClient,
	}
}
