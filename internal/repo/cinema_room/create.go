package cinemaroom_repo

import (
	"github.com/sale-tickets/manager-api/internal/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *cinemaRoomRepo) Create(req *view.CreateCinemaRoomReq) (uuid string, err error) {
	cinemaRoom := &model.CinemaRoom{
		MovieTheaterId: req.MovieTheaterId,
		Code:           req.Code,
	}

	err = r.db.Model(&model.CinemaRoom{}).Create(&cinemaRoom).Error
	if err != nil {
		return "", err
	}

	return cinemaRoom.Uuid, nil
}
