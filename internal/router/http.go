package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/godev-lib/golang/config"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type READY_HTTP struct{}

func HttpServer(config *config.Config) {
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}

		grpcAddr := fmt.Sprintf(
			"%s:%s",
			config.App.Host,
			config.App.GrpcPort,
		)
		manager_api.RegisterHealthHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
		manager_api.RegisterMovieTheaterHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
		manager_api.RegisterCinemaRoomServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
		manager_api.RegisterTheaterSeatingHandlerFromEndpoint(ctx, mux, grpcAddr, opts)

		httpAddr := fmt.Sprintf(
			":%s",
			config.App.HttpPort,
		)
		log.Printf("REST gateway running on %s", httpAddr)
		server := http.Server{
			Addr:    httpAddr,
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("error start http server: ", err.Error())
		}
	}()
}
