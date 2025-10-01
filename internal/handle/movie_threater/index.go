package moviethreater_controller

import (
	movietheater_service "github.com/sale-tickets/manager-api/internal/service/movie_theater"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

type movieTheaterController struct {
	movieTheaterService movietheater_service.MovietheaterService
	manager_api.UnimplementedMovieTheaterServer
}

func NewHandle() manager_api.MovieTheaterServer {
	return &movieTheaterController{
		movieTheaterService: movietheater_service.NewMovietheaterService(),
	}
}
