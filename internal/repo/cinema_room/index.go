package cinemaroom_repo

import (
	"github.com/sale-tickets/manager-api/internal/common/connection"
	"github.com/sale-tickets/manager-api/internal/view"

	"gorm.io/gorm"
)

type (
	cinemaRoomRepo struct {
		db *gorm.DB
	}
	CinemaRoomRepo interface {
		Create(req *view.CreateCinemaRoomReq) (uuid string, err error)
	}
)

func NewCinemaRoomRepo() CinemaRoomRepo {
	return &cinemaRoomRepo{
		db: connection.ConfigInfo.Database.GetConection(),
	}
}
