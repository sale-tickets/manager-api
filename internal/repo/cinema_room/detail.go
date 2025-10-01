package cinemaroom_repo

import (
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) Detail(req *view.DetailCinemaRoomReq) (*model.CinemaRoom, error) {
	cinemaRoom := &model.CinemaRoom{}

	queryStr := `
		SELECT
			uuid,
			movie_theater_id,
			code,
			created_at,
			updated_at
		FROM cinema_rooms
		WHERE uuid = ? AND deleted_at IS NULL
		LIMIT 1
		OFFSET 0
	`

	err := r.db.Raw(
		queryStr,
		req.Uuid,
	).Scan(&cinemaRoom).Error
	if err != nil {
		return nil, err
	}

	return cinemaRoom, nil
}
