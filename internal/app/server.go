package app

import (
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"
)

type UserServer struct {
	userService service.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserServer(userService service.UserService) *UserServer {
	return &UserServer{userService: userService}
}
