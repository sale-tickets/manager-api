package moviethreater_controller

import (
	"context"
	"errors"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"

	"google.golang.org/grpc/metadata"
)

func (c *movieTheaterController) Create(ctx context.Context, req *manager_api.CreateMovieTheaterReq) (*manager_api.CreateMovieTheaterRes, error) {
	reqData := view.CreateMovieTheaterReq{
		CreaterId:             utils.GetProfileId(ctx),
		CreateMovieTheaterReq: req,
	}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata invalid")
	}

	if len(md.Get("uuid")) == 0 {
		return nil, errors.New("uuid not found")
	}
	createId := md.Get("uuid")[0]
	reqData.CreaterId = createId

	id, err := c.movieTheaterService.Create(reqData)
	if err != nil {
		return nil, err
	}

	response := &manager_api.CreateMovieTheaterRes{
		Uuid:    id,
		Name:    reqData.Name,
		Address: reqData.Address,
	}
	return response, nil
}
