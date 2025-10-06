package theaterseating_repo

import (
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *theaterSeatingRepo) List(req *view.ListTheaterSeatingReq) ([]*model.TheaterSeating, error) {
	theaterSeatings := []*model.TheaterSeating{}

	queryStr := `
		SELECT
			*
		FROM theater_seatings
		WHERE
			cinema_room_id = ?
		LIMIT ?
		OFFSET ?
	`

	err := r.db.Raw(
		queryStr,
		req.Filter.CinemaRoomId,
		req.FilterBase.Limit,
		req.FilterBase.Offset,
	).Scan(&theaterSeatings).Error
	if err != nil {
		return nil, err
	}

	return theaterSeatings, nil
}
