package movietheater_repo

import (
	"time"

	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *movietheaterRepo) Update(req view.UpdateMovieTheaterReq) error {
	queryStr := `
		UPDATE movie_theaters
		SET
			name = ?,
			address = ?,
			updated_at = ?
		WHERE 
			uuid = ? 
			AND creater_id = ?
			AND deleted_at IS NULL
	`

	err := r.db.Exec(
		queryStr,
		req.Name,
		req.Address,
		time.Now(),
		req.Uuid,
		req.CreaterId,
	).Error
	if err != nil {
		return err
	}

	return nil
}
