package theaterseating_repo

import (
	"github.com/sale-tickets/golang-common/model"
	"github.com/sale-tickets/manager-api/internal/view"
	"gorm.io/gorm"
)

type (
	theaterSeatingRepo struct {
		db *gorm.DB
	}
	TheaterSeatingRepo interface {
		Create(req *view.CreateTheaterSeatingReq) ([]string, error)
		List(req *view.ListTheaterSeatingReq) ([]*model.TheaterSeating, error)
		Count(req *view.ListTheaterSeatingReq) (int64, error)
	}
)

func NewTheaterSeatingRepo(db *gorm.DB) TheaterSeatingRepo {
	return &theaterSeatingRepo{
		db: db,
	}
}
