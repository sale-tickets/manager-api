package movietheater_service

import (
	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (s *movieTheaterService) Detail(req view.DetailMovieTheaterReq) (*manager_api.DetailMovieTheaterRes, error) {
	movieTheater, err := s.movieTheaterRepo.Detail(req)
	if err != nil {
		return nil, err
	}

	result := &manager_api.DetailMovieTheaterRes{
		Uuid:      movieTheater.Uuid,
		Name:      movieTheater.Name,
		Address:   movieTheater.Address,
		CreaterId: movieTheater.CreaterId,
		CreatedAt: movieTheater.CreatedAt.String(),
		UpdatedAt: movieTheater.UpdatedAt.String(),
	}

	return result, nil
}
