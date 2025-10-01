package cinemaroom_repo

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) CountList(req *view.ListCinemaRoomReq) (int32, error) {
	var count int32

	queryStr := fmt.Sprintf(`
			SELECT count(uuid)
			FROM cinema_rooms
			WHERE %s AND deleted_at IS NULL
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("code"),
		),
	)

	err := r.db.Raw(
		queryStr,
		utils.AddLikeValue(req.Filter.Code),
	).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
