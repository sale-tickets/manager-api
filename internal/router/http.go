package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sale-tickets/manager-api/internal/common/connection"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func HttpServer(startedGrpc <-chan bool, errStartHttpServer chan<- error) {
	<-startedGrpc

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	grpcAddr := fmt.Sprintf(
		"%s:%s",
		connection.ConfigInfo.App.Host,
		connection.ConfigInfo.App.GrpcPort,
	)
	manager_api.RegisterHealthHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	manager_api.RegisterMovieTheaterHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	manager_api.RegisterCinemaRoomServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)

	httpAddr := fmt.Sprintf(
		":%s",
		connection.ConfigInfo.App.HttpPort,
	)
	log.Printf("REST gateway running on %s", httpAddr)
	server := http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		errStartHttpServer <- err
	}
}
