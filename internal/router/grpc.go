package router

import (
	"fmt"
	"log"
	"net"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"

	"github.com/sale-tickets/manager-api/internal/common/connection"
	"github.com/sale-tickets/manager-api/internal/common/middleware"
	cinemaroom_controller "github.com/sale-tickets/manager-api/internal/handle/cinema_room"
	health_controller "github.com/sale-tickets/manager-api/internal/handle/health"
	moviethreater_controller "github.com/sale-tickets/manager-api/internal/handle/movie_threater"

	"google.golang.org/grpc"
)

func GrpcServer(startedGrpc chan<- bool, errStartGrpcServer chan<- error) {
	port := fmt.Sprintf(":%s", connection.ConfigInfo.App.GrpcPort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.ValidateToken,
			middleware.GetProfileId,
		),
	)
	manager_api.RegisterHealthServer(s, health_controller.NewHandle())
	manager_api.RegisterMovieTheaterServer(s, moviethreater_controller.NewHandle())
	manager_api.RegisterCinemaRoomServiceServer(s, cinemaroom_controller.NewHandle())

	log.Printf("gRPC server running on %s", port)
	startedGrpc <- true
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		errStartGrpcServer <- err
	}
}
