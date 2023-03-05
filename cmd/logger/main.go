package main

import (
	"net"
	"sirawit/shop/internal/app"
	"sirawit/shop/internal/config"
	"sirawit/shop/pkg/pb"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// load config

	config, err := config.LoadLoggerConfig(".")
	if err != nil {
		log.Err(err).Msg("cannot load config file")
	}

	server := app.NewLoggerServer()
	if err != nil {
		log.Fatal().Msg("cannot create logger server")
	}

	// grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GrpcLoggerServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot create listener")
	}
	log.Printf("start gRPC server at %v", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("cannot start grpc server")
	}
}
