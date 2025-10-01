package movietheater_repo

import (
	"github.com/sale-tickets/manager-api/internal/common/connection"
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"

	"gorm.io/gorm"
)

type (
	movietheaterRepo struct {
		db *gorm.DB
	}
	MovietheaterRepo interface {
		Create(req view.CreateMovieTheaterReq) (id string, err error)
		Update(req view.UpdateMovieTheaterReq) error
		GetList(req view.GetListMovieTheaterReq) ([]model.MovieTheater, error)
		CountList(req view.GetListMovieTheaterReq) (int32, error)
		Detail(req view.DetailMovieTheaterReq) (*model.MovieTheater, error)
	}
)

func NewMovietheaterRepo() MovietheaterRepo {
	return &movietheaterRepo{
		db: connection.ConfigInfo.Database.GetConection(),
	}
}
