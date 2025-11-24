package cinemaroom_repo

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) CountList(req *view.ListCinemaRoomReq) (int32, error) {
	var count int32

	queryStr := fmt.Sprintf(`
			SELECT count(cr.uuid)
			FROM cinema_rooms AS cr
			JOIN movie_theaters AS mt ON mt.uuid = cr.movie_theater_id
			WHERE 
				%s 
				AND mt.creater_id = ?
				AND cr.deleted_at IS NULL
				AND mt.deleted_at IS NULL
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("cr.movie_theater_id"),
		),
	)

	err := r.db.Raw(
		queryStr,
		utils.AddLikeValue(req.Filter.Code),
		req.CreateId,
	).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
