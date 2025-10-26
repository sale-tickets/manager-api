package movie_controller

import (
	"context"

	"github.com/godev-lib/golang/orm"
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (c *movieController) List(ctx context.Context, req *manager_api.ListMovieReq) (*manager_api.ListMovieRes, error) {
	films, err := c.repo.Find(orm.Filter{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	result := &manager_api.ListMovieRes{}
	for _, item := range films {
		result.Data = append(result.Data, &manager_api.MovieModel{
			Uuid:        item.Uuid,
			Name:        item.Name,
			CategoryId:  item.CategoryID,
			Description: item.Description,
			Thumbnail:   item.Thumbnail,
			CreatedAt:   item.CreatedAt.String(),
			UpdatedAt:   item.UpdatedAt.String(),
		})
	}
	return result, nil
}
