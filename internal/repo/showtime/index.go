package showtime_repo

import (
	"github.com/godev-lib/golang/orm"
	"github.com/sale-tickets/golang-common/model"
	"gorm.io/gorm"
)

type ShowtimeRepo struct {
	orm.DataMethod[model.Showtime]
}

func NewShowtimeRepo(db *gorm.DB) *ShowtimeRepo {
	return &ShowtimeRepo{
		orm.NewOrm[model.Showtime](db),
	}
}
