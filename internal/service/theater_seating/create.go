package theaterseating_service

import "github.com/sale-tickets/manager-api/internal/view"

func (s *theaterSeatingService) Create(req *view.CreateTheaterSeatingReq) ([]string, error) {
	return s.repo.Create(req)
}
