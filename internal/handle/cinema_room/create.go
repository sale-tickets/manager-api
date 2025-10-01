package cinemaroom_controller

import (
	"context"

	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (c *cinemaRoomController) Create(ctx context.Context, req *manager_api.CreateCinemaRoomReq) (*manager_api.CreateCinemaRoomRes, error) {
	reqDate := &view.CreateCinemaRoomReq{CreateCinemaRoomReq: req}
	if err := reqDate.Validate(); err != nil {
		return nil, err
	}

	uuid, err := c.service.Create(reqDate)
	if err != nil {
		return nil, err
	}

	result := &manager_api.CreateCinemaRoomRes{
		Uuid:           uuid,
		Code:           req.Code,
		MovieTheaterId: req.MovieTheaterId,
	}
	return result, nil
}
