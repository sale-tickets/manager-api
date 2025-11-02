package theaterseating_repo

import (
	"github.com/google/uuid"

	"github.com/sale-tickets/golang-common/model"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (r *theaterSeatingRepo) Create(req *view.CreateTheaterSeatingReq) ([]string, error) {
	theaterSeatings := []*model.TheaterSeating{}
	ids := []string{}

	for _, item := range req.Theaters {
		uuidBuf, err := uuid.NewV7()
		if err != nil {
			return nil, err
		}
		theaterSeatings = append(theaterSeatings, &model.TheaterSeating{
			Base: model.Base{
				Uuid: uuidBuf.String(),
			},
			CinemaRoomId: req.CinemaRoomId,
			Name:         item.Name,
			RowIndex:     item.RowIndex,
			ColIndex:     item.ColIndex,
			Code:         item.Code,
		})
		ids = append(ids, uuidBuf.String())
	}

	err := r.db.
		Model(&model.TheaterSeating{}).
		CreateInBatches(&theaterSeatings, len(theaterSeatings)).
		Error
	if err != nil {
		return nil, err
	}

	return ids, nil
}
