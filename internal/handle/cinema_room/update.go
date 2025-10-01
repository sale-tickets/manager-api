package cinemaroom_controller

import (
	"context"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (c *cinemaRoomController) Update(ctx context.Context, req *manager_api.UpdateCinemaRoomReq) (*manager_api.UpdateCinemaRoomRes, error) {
	reqData := &view.UpdateCinemaRoomReq{UpdateCinemaRoomReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	err := c.service.Update(reqData)
	if err != nil {
		return nil, err
	}

	result := &manager_api.UpdateCinemaRoomRes{}

	return result, nil
}
