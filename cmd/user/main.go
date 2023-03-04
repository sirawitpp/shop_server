package main

import (
	"context"
	"net/http"
	"sirawit/shop/internal/app"
	"sirawit/shop/internal/repository"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/config"
	"sirawit/shop/pkg/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
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
	// setup server&service

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userServer := app.NewUserServer(config, userSvc)

	// server options

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// start server

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userServer)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot reigster server")
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	srv := &http.Server{
		Addr:    config.HttpServerAddress,
		Handler: mux,
	}
	log.Info().Msgf("start server at %v", config.HttpServerAddress)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("cannot start  server")
	}
}
