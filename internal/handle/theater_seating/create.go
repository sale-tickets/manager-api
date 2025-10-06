package theaterseating_controller

import (
	"context"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (c *theaterSeatingController) CreateTheaterSeating(ctx context.Context, req *manager_api.CreateTheaterSeatingReq) (*manager_api.CreateTheaterSeatingRes, error) {
	reqData := &view.CreateTheaterSeatingReq{CreateTheaterSeatingReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	ids, err := c.service.Create(reqData)
	if err != nil {
		return nil, err
	}

	result := &manager_api.CreateTheaterSeatingRes{
		ListId: ids,
	}
	return result, nil
}
