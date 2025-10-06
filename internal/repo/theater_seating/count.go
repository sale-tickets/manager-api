package theaterseating_repo

import "github.com/sale-tickets/manager-api/internal/view"

func (r *theaterSeatingRepo) Count(req *view.ListTheaterSeatingReq) (int64, error) {
	var count int64

	queryStr := `
		SELECT
			count(uuid)
		FROM theater_seatings
		WHERE
			cinema_room_id = ?
	`

	err := r.db.Raw(
		queryStr,
		req.Filter.CinemaRoomId,
	).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
