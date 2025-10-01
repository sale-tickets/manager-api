package cinemaroom_controller

import (
	"context"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (c *cinemaRoomController) Detail(ctx context.Context, req *manager_api.DetailCinemaRoomReq) (*manager_api.DetailCinemaRoomRes, error) {
	reqData := &view.DetailCinemaRoomReq{DetailCinemaRoomReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}
	return c.service.Detail(reqData)
}
