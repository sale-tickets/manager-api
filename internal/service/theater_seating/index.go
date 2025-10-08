package theaterseating_service

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	theaterseating_repo "github.com/sale-tickets/manager-api/internal/repo/theater_seating"
	"github.com/sale-tickets/manager-api/internal/view"
)

type (
	theaterSeatingService struct {
		repo theaterseating_repo.TheaterSeatingRepo
	}
	TheaterSeatingService interface {
		Create(req *view.CreateTheaterSeatingReq) ([]string, error)
		List(req *view.ListTheaterSeatingReq) (*manager_api.ListTheaterSeatingRes, error)
	}
)

func NewTheaterSeatingService(repo theaterseating_repo.TheaterSeatingRepo) TheaterSeatingService {
	return &theaterSeatingService{
		repo: repo,
	}
}
