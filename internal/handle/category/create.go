package category

import (
	"context"

	"github.com/google/uuid"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/golang-common/model"
)

func (h *CategoryHandle) Create(ctx context.Context, req *manager_api.CreateCategoryReq) (*manager_api.CreateCategoryRes, error) {
	uuidBuf, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	category, err := h.repo.CreateOne(&model.Category{
		Base: model.Base{
			Uuid: uuidBuf.String(),
		},
		Name: req.Data.Name,
		Code: req.Data.Code,
	})
	if err != nil {
		return nil, err
	}

	result := &manager_api.CreateCategoryRes{
		Data: &manager_api.CategoryModel{
			Uuid:      category.Uuid,
			Name:      category.Name,
			Code:      category.Code,
			CreatedAt: category.CreatedAt.String(),
			UpdatedAt: category.UpdatedAt.String(),
		},
	}
	return result, nil
}
