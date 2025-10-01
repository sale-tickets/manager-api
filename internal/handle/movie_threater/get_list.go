package moviethreater_controller

import (
	"context"

	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/duyhung2k4/sale-tickets-golang-common/manager-api/proto"
)

func (c *movieTheaterController) GetList(ctx context.Context, req *manager_api.GetListMovieTheaterReq) (*manager_api.GetListMovieTheaterRes, error) {
	reqData := view.GetListMovieTheaterReq{GetListMovieTheaterReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	result, err := c.movieTheaterService.GetList(reqData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
