package cinemaroom_repo

import (
	"fmt"

	"github.com/sale-tickets/golang-common/model"
	"github.com/sale-tickets/manager-api/internal/common/utils"

	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) List(req *view.ListCinemaRoomReq) ([]*model.CinemaRoom, error) {
	list := []*model.CinemaRoom{}

	queryStr := fmt.Sprintf(`
			SELECT
				cr.uuid,
				cr.movie_theater_id,
				cr.code,
				cr.created_at,
				cr.updated_at
			FROM cinema_rooms AS cr
			JOIN movie_theaters AS mt ON mt.uuid = cr.movie_theater_id
			WHERE
				%s
				AND mt.creater_id = ?
				AND cr.deleted_at IS NULL
				AND mt.deleted_at IS NULL
			ORDER BY code ASC
			LIMIT ?
			OFFSET ?
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("cr.code"),
		),
	)

	err := r.db.Raw(
		queryStr,
		utils.AddLikeValue(req.Filter.Code),
		req.CreateId,
		req.FilterBase.Limit,
		req.FilterBase.Offset,
	).Scan(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
