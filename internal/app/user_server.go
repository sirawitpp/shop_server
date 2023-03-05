package app

import (
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"

	"google.golang.org/grpc"
)

type UserServer struct {
	userService service.UserService
	pb.UnimplementedUserServiceServer
	conn         *grpc.ClientConn
	loggerClient pb.LoggerServiceClient
}

func NewUserServer(userService service.UserService, conn *grpc.ClientConn) *UserServer {
	loggerClient := pb.NewLoggerServiceClient(conn)
	return &UserServer{
		userService:  userService,
		conn:         conn,
		loggerClient: loggerClient,
	}
}
