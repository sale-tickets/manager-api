package movie_controller

import (
	"context"

	"github.com/google/uuid"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (c *movieController) Create(ctx context.Context, req *manager_api.CreateMovieReq) (*manager_api.CreateMovieRes, error) {
	reqData := view.CreateMovieReq{
		CreateMovieReq: req,
	}

	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	uuidBuf, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	movie := &model.Movie{
		Base: model.Base{
			Uuid: uuidBuf.String(),
		},
		Name:        reqData.Data.Name,
		CategoryID:  req.Data.CategoryId,
		Description: reqData.Data.Description,
		Thumbnail:   reqData.Data.Thumbnail,
	}

	newMovie, err := c.repo.CreateOne(movie)
	if err != nil {
		return nil, err
	}

	result := &manager_api.CreateMovieRes{
		Data: &manager_api.MovieModel{
			Uuid:        newMovie.Uuid,
			Name:        newMovie.Name,
			CategoryId:  req.Data.CategoryId,
			Description: newMovie.Description,
			Thumbnail:   newMovie.Thumbnail,
		},
	}
	return result, nil
}
