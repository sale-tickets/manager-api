package moviethreater_controller

import (
	"context"

	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/duyhung2k4/sale-tickets-golang-common/manager-api/proto"
)

func (c *movieTheaterController) Update(ctx context.Context, req *manager_api.UpdateMovieTheaterReq) (*manager_api.UpdateMovieTheaterRes, error) {
	reqData := view.UpdateMovieTheaterReq{UpdateMovieTheaterReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	result, err := c.movieTheaterService.Update(reqData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
