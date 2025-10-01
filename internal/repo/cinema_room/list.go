package cinemaroom_repo

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) List(req *view.ListCinemaRoomReq) ([]*model.CinemaRoom, error) {
	list := []*model.CinemaRoom{}

	queryStr := fmt.Sprintf(`
			SELECT
				uuid,
				movie_theater_id,
				code,
				created_at,
				updated_at
			FROM cinema_rooms
			WHERE
				%s
				AND deleted_at IS NULL
			ORDER BY code ASC
			LIMIT ?
			OFFSET ?
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("code"),
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
