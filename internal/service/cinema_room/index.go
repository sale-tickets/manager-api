package cinemaroom_service

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	cinemaroom_repo "github.com/sale-tickets/manager-api/internal/repo/cinema_room"
	"github.com/sale-tickets/manager-api/internal/view"
)

type (
	cinemaRoomService struct {
		repo cinemaroom_repo.CinemaRoomRepo
	}
	CinemaRoomService interface {
		Create(req *view.CreateCinemaRoomReq) (uuid string, err error)
		Update(req *view.UpdateCinemaRoomReq) error
		Detail(req *view.DetailCinemaRoomReq) (*manager_api.DetailCinemaRoomRes, error)
		List(req *view.ListCinemaRoomReq) (*manager_api.ListCinemaRoomRes, error)
	}
)

func NewCinemaRoomServic() CinemaRoomService {
	return &cinemaRoomService{
		repo: cinemaroom_repo.NewCinemaRoomRepo(),
	}
}
