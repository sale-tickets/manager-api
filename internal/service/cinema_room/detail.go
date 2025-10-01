package cinemaroom_service

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/manager-api/internal/view"
)

func (s *cinemaRoomService) Detail(req *view.DetailCinemaRoomReq) (*manager_api.DetailCinemaRoomRes, error) {
	cinemaRoom, err := s.repo.Detail(req)
	if err != nil {
		return nil, err
	}

	result := &manager_api.DetailCinemaRoomRes{
		Uuid:           cinemaRoom.Uuid,
		MovieTheaterId: cinemaRoom.MovieTheaterId,
		Code:           cinemaRoom.Code,
		CreatedAt:      cinemaRoom.CreatedAt.String(),
		UpdatedAt:      cinemaRoom.UpdatedAt.String(),
	}

	return result, nil
}
