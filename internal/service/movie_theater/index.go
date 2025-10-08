package movietheater_service

import (
	movietheater_repo "github.com/sale-tickets/manager-api/internal/repo/movie_theater"
	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

type (
	movieTheaterService struct {
		movieTheaterRepo movietheater_repo.MovietheaterRepo
	}
	MovietheaterService interface {
		Create(req view.CreateMovieTheaterReq) (id string, err error)
		Update(req view.UpdateMovieTheaterReq) (*manager_api.UpdateMovieTheaterRes, error)
		GetList(req view.GetListMovieTheaterReq) (*manager_api.GetListMovieTheaterRes, error)
		Detail(req view.DetailMovieTheaterReq) (*manager_api.DetailMovieTheaterRes, error)
	}
)

func NewMovietheaterService(repo movietheater_repo.MovietheaterRepo) MovietheaterService {
	return &movieTheaterService{
		movieTheaterRepo: repo,
	}
}
