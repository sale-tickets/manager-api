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
			WHERE
				%s
				AND cr.deleted_at IS NULL
			ORDER BY code ASC
			LIMIT ?
			OFFSET ?
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("cr.movie_theater_id"),
		),
	)

	err := r.db.Raw(
		queryStr,
		utils.AddLikeValue(req.Filter.Code),
		req.FilterBase.Limit,
		req.FilterBase.Offset,
	).Scan(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
