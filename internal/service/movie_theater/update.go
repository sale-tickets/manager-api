package movietheater_service

import (
	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (s *movieTheaterService) Update(req view.UpdateMovieTheaterReq) (*manager_api.UpdateMovieTheaterRes, error) {
	err := s.movieTheaterRepo.Update(req)
	if err != nil {
		return nil, err
	}

	result := &manager_api.UpdateMovieTheaterRes{}
	return result, nil
}
