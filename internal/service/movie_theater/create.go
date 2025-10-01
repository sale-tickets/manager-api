package movietheater_service

import "github.com/sale-tickets/manager-api/internal/view"

func (s *movieTheaterService) Create(req view.CreateMovieTheaterReq) (id string, err error) {
	return s.movieTheaterRepo.Create(req)
}
