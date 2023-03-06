package main

import (
	"context"
	"net/http"
	"sirawit/shop/internal/app"
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/repository"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	// load config

	config, err := config.LoadUserConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	// intdb

	db, err := repository.ConnectToUserDB(config.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	log.Info().Msg("connect to user db!")

	//grpc client
	log.Info().Msg("try to connect to logger service")
	conn, err := grpc.Dial(config.GrpcLoggerServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connet to logger service")
	}
	defer conn.Close()
	log.Info().Msgf("start grpc client at %v", config.GrpcLoggerServerAddress)

	// setup server&service

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo, config)
	userServer := app.NewUserServer(userSvc, conn)

	// server options

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// setup server

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userServer)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot reigster server")
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	handler := app.HttpLogger(mux)
	srv := &http.Server{
		Addr:    config.HttpServerAddress,
		Handler: handler,
	}

	//start server

	log.Info().Msgf("start http server at %v", config.HttpServerAddress)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("cannot start  server")
	}
}
