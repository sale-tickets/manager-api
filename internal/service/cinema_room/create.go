package cinemaroom_service

import "github.com/sale-tickets/manager-api/internal/view"

func (s *cinemaRoomService) Create(req *view.CreateCinemaRoomReq) (uuid string, err error) {
	return s.repo.Create(req)
}
