package category

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	category_repo "github.com/sale-tickets/manager-api/internal/repo/category"
)

type CategoryHandle struct {
	repo *category_repo.CategoryRepo
	manager_api.UnimplementedCategoryServer
}

func NewCategoryHandle(repo *category_repo.CategoryRepo) manager_api.CategoryServer {
	return &CategoryHandle{
		repo: repo,
	}
}
