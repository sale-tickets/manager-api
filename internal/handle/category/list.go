package category

import (
	"context"

	"github.com/godev-lib/golang/orm"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (h *CategoryHandle) List(ctx context.Context, req *manager_api.ListCategoryReq) (*manager_api.ListCategoryRes, error) {
	categories, err := h.repo.Find(orm.Filter{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	result := &manager_api.ListCategoryRes{}
	for _, item := range categories {
		result.Data = append(result.Data, &manager_api.CategoryModel{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Code:      item.Code,
			CreatedAt: item.CreatedAt.String(),
			UpdatedAt: item.UpdatedAt.String(),
		})
	}
	return result, nil
}
