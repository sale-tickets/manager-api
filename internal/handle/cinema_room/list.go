package cinemaroom_controller

import (
	"context"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (c *cinemaRoomController) List(ctx context.Context, req *manager_api.ListCinemaRoomReq) (*manager_api.ListCinemaRoomRes, error) {
	reqData := &view.ListCinemaRoomReq{
		CreateId:          utils.GetProfileId(ctx),
		ListCinemaRoomReq: req,
	}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}
	return c.service.List(reqData)
}
