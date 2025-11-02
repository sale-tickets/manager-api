package cinemaroom_controller

import (
	"context"

	"github.com/sale-tickets/manager-api/internal/common/utils"
	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (c *cinemaRoomController) Create(ctx context.Context, req *manager_api.CreateCinemaRoomReq) (*manager_api.CreateCinemaRoomRes, error) {
	reqDate := &view.CreateCinemaRoomReq{
		CreaterId:           utils.GetProfileId(ctx),
		CreateCinemaRoomReq: req,
	}
	if err := reqDate.Validate(); err != nil {
		return nil, err
	}

	uuid, err := c.service.Create(reqDate)
	if err != nil {
		return nil, err
	}

	seat := []string{
		"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "A9", "A10",
		"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9", "B10",
		"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10",
		"D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "D10",
		"E1", "E2", "E3", "E4", "E5", "E6", "E7", "E8", "E9", "E10",
		"F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10",
	}
	seatings := []*manager_api.TheaterSeatingModel{}
	for _, code := range seat {
		seatings = append(seatings, &manager_api.TheaterSeatingModel{
			CinemaRoomId: uuid,
			Code:         code,
		})
	}
	_, err = c.theaterSeatingService.Create(&view.CreateTheaterSeatingReq{
		CreateTheaterSeatingReq: &manager_api.CreateTheaterSeatingReq{
			CinemaRoomId: uuid,
			Theaters:     seatings,
		},
	})
	if err != nil {
		return nil, err
	}

	result := &manager_api.CreateCinemaRoomRes{
		Uuid:           uuid,
		Code:           req.Code,
		MovieTheaterId: req.MovieTheaterId,
	}
	return result, nil
}
