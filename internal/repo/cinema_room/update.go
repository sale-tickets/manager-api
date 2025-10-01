package cinemaroom_repo

import (
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) Update(req *view.UpdateCinemaRoomReq) error {
	err := r.db.
		Model(&model.CinemaRoom{}).
		Where("uuid = ?", req.Uuid).
		Update("code = ?", req.Code).Error
	return err
}
