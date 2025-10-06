package theaterseating_controller

import (
	"context"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (c *theaterSeatingController) ListTheaterSeating(ctx context.Context, req *manager_api.ListTheaterSeatingReq) (*manager_api.ListTheaterSeatingRes, error) {
	reqData := &view.ListTheaterSeatingReq{ListTheaterSeatingReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}
	return c.service.List(reqData)
}
