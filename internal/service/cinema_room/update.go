package cinemaroom_service

import "github.com/sale-tickets/manager-api/internal/view"

func (s *cinemaRoomService) Update(req *view.UpdateCinemaRoomReq) error {
	return s.repo.Update(req)
}
