package movietheater_repo

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *movietheaterRepo) CountList(req view.GetListMovieTheaterReq) (int32, error) {
	var count int32
	queryStr := fmt.Sprintf(`
		SELECT count(uuid)
		FROM movie_theaters
		WHERE 
			%s 
			AND creater_id = ?
			AND deleted_at IS NULL
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("name"),
			utils.AddLikeClause("address"),
		),
	)
	err := r.db.Raw(
		queryStr,
		req.CreaterId,
		utils.AddLikeValue(req.Filter.Name),
		utils.AddLikeValue(req.Filter.Address),
	).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
