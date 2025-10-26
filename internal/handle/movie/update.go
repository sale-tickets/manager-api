package movie_controller

import (
	"context"

	"github.com/godev-lib/golang/orm"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/golang-common/model"

	"github.com/sale-tickets/manager-api/internal/view"
)

func (c *movieController) Update(ctx context.Context, req *manager_api.UpdateMovieReq) (*manager_api.UpdateMovieRes, error) {
	reqData := view.UpdateMovieReq{
		UpdateMovieReq: req,
	}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	dataUpdate := &model.Movie{
		CategoryID:  req.Data.CategoryId,
		Name:        req.Data.Name,
		Description: req.Data.Description,
		Thumbnail:   req.Data.Thumbnail,
	}

	newMovie, err := c.repo.Update(dataUpdate, orm.Filter{
		IsReturning: true,
		Conditions: []orm.Condition{
			{
				Query: "uuid = ?",
				Arg:   reqData.Id,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	result := &manager_api.UpdateMovieRes{
		Data: &manager_api.MovieModel{
			Uuid:        newMovie.Uuid,
			CategoryId:  newMovie.CategoryID,
			Name:        newMovie.Name,
			Description: newMovie.Description,
			Thumbnail:   newMovie.Thumbnail,
			CreatedAt:   newMovie.CreatedAt.String(),
			UpdatedAt:   newMovie.UpdatedAt.String(),
		},
	}

	return result, nil
}
