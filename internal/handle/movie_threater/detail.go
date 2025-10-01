package moviethreater_controller

import (
	"context"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (c *movieTheaterController) Detail(ctx context.Context, req *manager_api.DetailMovieTheaterReq) (*manager_api.DetailMovieTheaterRes, error) {
	reqData := view.DetailMovieTheaterReq{
		CreaterId:             utils.GetProfileId(ctx),
		DetailMovieTheaterReq: req,
	}

	result, err := c.movieTheaterService.Detail(reqData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
