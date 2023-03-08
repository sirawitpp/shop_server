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
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// load config

	config, err := config.LoadProductConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	// intdb

	db, err := repository.ConnectToProductDB(config.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	log.Info().Msg("connect to product db!")

	//setup service && server

	productQuery := repository.NewProductQuery(db)
	productService := service.NewProductService(productQuery, config)
	server := app.NewProductServer(productService, config)

	//json options

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
	err = pb.RegisterProductServiceHandlerServer(ctx, grpcMux, server)
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
