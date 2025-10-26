package movie_controller

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	movie_repo "github.com/sale-tickets/manager-api/internal/repo/movie"
)

type (
	movieController struct {
		repo *movie_repo.MovieRepo
		manager_api.UnimplementedMovieServer
	}
)

func NewHandle(repo *movie_repo.MovieRepo) manager_api.MovieServer {
	return &movieController{
		repo: repo,
	}
}
