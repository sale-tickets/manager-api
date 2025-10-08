package cinemaroom_repo

import (
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"

	"gorm.io/gorm"
)

type (
	cinemaRoomRepo struct {
		db *gorm.DB
	}
	CinemaRoomRepo interface {
		Create(req *view.CreateCinemaRoomReq) (id string, err error)
		Update(req *view.UpdateCinemaRoomReq) error
		List(req *view.ListCinemaRoomReq) ([]*model.CinemaRoom, error)
		CountList(req *view.ListCinemaRoomReq) (int32, error)
		Detail(req *view.DetailCinemaRoomReq) (*model.CinemaRoom, error)
	}
)

func NewCinemaRoomRepo(db *gorm.DB) CinemaRoomRepo {
	return &cinemaRoomRepo{
		db: db,
	}
}
