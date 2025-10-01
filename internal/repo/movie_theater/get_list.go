package movietheater_repo

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *movietheaterRepo) GetList(req view.GetListMovieTheaterReq) ([]model.MovieTheater, error) {
	results := []model.MovieTheater{}

	queryStr := fmt.Sprintf(`
		SELECT
			uuid,
			creater_id,
			name,
			address
		FROM movie_theaters
		WHERE
			%s
			AND deleted_at IS NULL
		LIMIT ?
		OFFSET ?
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("name"),
			utils.AddLikeClause("address"),
		),
	)

	err := r.db.Raw(
		queryStr,
		utils.AddLikeValue(req.Filter.Name),
		utils.AddLikeValue(req.Filter.Address),
		req.FilterBase.Limit,
		req.FilterBase.Offset,
	).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
