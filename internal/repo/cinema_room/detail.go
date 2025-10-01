package cinemaroom_repo

import (
	"errors"

	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) Detail(req *view.DetailCinemaRoomReq) (*model.CinemaRoom, error) {
	cinemaRoom := &model.CinemaRoom{}

	queryStr := `
		SELECT
			cr.uuid,
			cr.movie_theater_id,
			cr.code,
			cr.created_at,
			cr.updated_at
		FROM cinema_rooms AS cr
		JOIN movie_theaters AS mt ON mt.uuid = cr.movie_theater_id
		WHERE 
			cr.uuid = ? 
			AND mt.creater_id = ?
			AND cr.deleted_at IS NULL
			AND mt.deleted_at IS NULL
		LIMIT 1
		OFFSET 0
	`

	err := r.db.Raw(
		queryStr,
		req.Uuid,
		req.CreaterId,
	).Scan(&cinemaRoom).Error
	if err != nil {
		return nil, err
	}

	if cinemaRoom.Uuid == "" {
		return nil, errors.New("cinema not found")
	}

	return cinemaRoom, nil
}
