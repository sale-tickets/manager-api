package movie_repo

import (
	"github.com/godev-lib/golang/orm"
	"github.com/sale-tickets/manager-api/internal/model"
	"gorm.io/gorm"
)

type MovieRepo struct {
	orm.DataMethod[model.Movie]
}

func NewMovieRepo(db *gorm.DB) *MovieRepo {
	return &MovieRepo{
		orm.NewOrm[model.Movie](db),
	}
}
