package app

import (
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/config"
	"sirawit/shop/pkg/pb"
	"sirawit/shop/pkg/token"
)

type UserServer struct {
	config      config.UserConfig
	userService service.UserService
	pb.UnimplementedUserServiceServer
	tokenManager token.Manager
}

func NewUserServer(
	config config.UserConfig,
	userService service.UserService,
) *UserServer {
	tokenManager := token.NewJWTManager(config.Sign)
	return &UserServer{
		userService:  userService,
		config:       config,
		tokenManager: tokenManager,
	}
}
