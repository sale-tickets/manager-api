package category_repo

import (
	"github.com/godev-lib/golang/orm"
	"github.com/sale-tickets/golang-common/model"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	orm.DataMethod[model.Category]
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		orm.NewOrm[model.Category](db),
	}
}
