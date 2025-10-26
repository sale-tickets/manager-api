package cinemaroom_repo

import (
	"github.com/google/uuid"
	"github.com/sale-tickets/golang-common/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) Create(req *view.CreateCinemaRoomReq) (id string, err error) {
	uuidBuf, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	cinemaRoom := &model.CinemaRoom{
		Base: model.Base{
			Uuid: uuidBuf.String(),
		},
		MovieTheaterId: req.MovieTheaterId,
		Code:           req.Code,
	}

	err = r.db.Model(&model.CinemaRoom{}).Create(&cinemaRoom).Error
	if err != nil {
		return "", err
	}

	return cinemaRoom.Uuid, nil
}
