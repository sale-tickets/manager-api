package router

import (
	"fmt"
	"log"
	"net"

	"github.com/godev-lib/golang/config"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"

	"github.com/sale-tickets/manager-api/internal/common/middleware"

	"google.golang.org/grpc"
)

type READY_GRPC struct{}

func GrpcServer(
	config *config.Config,
	healthServer manager_api.HealthServer,
	movieTheaterServer manager_api.MovieTheaterServer,
	cinemaRoomServiceServer manager_api.CinemaRoomServiceServer,
	theaterSeatingServer manager_api.TheaterSeatingServer,
	movieServer manager_api.MovieServer,
	showtimeServer manager_api.ShowtimeServer,
) {
	go func() {
		port := fmt.Sprintf(":%s", config.App.GrpcPort)
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
		manager_api.RegisterHealthServer(s, healthServer)
		manager_api.RegisterMovieTheaterServer(s, movieTheaterServer)
		manager_api.RegisterCinemaRoomServiceServer(s, cinemaRoomServiceServer)
		manager_api.RegisterTheaterSeatingServer(s, theaterSeatingServer)
		manager_api.RegisterMovieServer(s, movieServer)
		manager_api.RegisterShowtimeServer(s, showtimeServer)

		log.Printf("gRPC server running on %s", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalln("error start grpc server: ", err.Error())
		}
	}()
}
