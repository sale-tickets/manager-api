package movietheater_repo

import (
	"errors"

	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *movietheaterRepo) Detail(req view.DetailMovieTheaterReq) (*model.MovieTheater, error) {
	movieTheater := &model.MovieTheater{}

	queryStr := `
		SELECT
			uuid,
			name,
			address,
			creater_id,
			created_at,
			updated_at
		FROM movie_theaters
		WHERE 
			uuid = ?
			AND creater_id = ?
			AND deleted_at IS NULL
		LIMIT 1
	`

	err := r.db.Raw(
		queryStr,
		req.Uuid,
		req.CreaterId,
	).Scan(&movieTheater).Error
	if err != nil {
		return nil, err
	}

	if movieTheater.Uuid == "" {
		return nil, errors.New("movie theater not found")
	}

	return movieTheater, nil
}
