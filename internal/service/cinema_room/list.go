package cinemaroom_service

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/golang-common/model"

	"github.com/sale-tickets/manager-api/internal/view"
)

func (s *cinemaRoomService) List(req *view.ListCinemaRoomReq) (*manager_api.ListCinemaRoomRes, error) {
	var chanErrCount, chanErrList = make(chan error, 1), make(chan error, 1)
	var chanCount = make(chan int32, 1)
	var chanList = make(chan []*model.CinemaRoom, 1)

	go func() {
		count, err := s.repo.CountList(req)
		chanCount <- count
		chanErrCount <- err
	}()

	go func() {
		list, err := s.repo.List(req)
		chanList <- list
		chanErrList <- err
	}()

	if err := <-chanErrCount; err != nil {
		return nil, err
	}
	if err := <-chanErrList; err != nil {
		return nil, err
	}

	result := &manager_api.ListCinemaRoomRes{
		Total: <-chanCount,
		Data:  make([]*manager_api.CinemaRoomModel, 0),
	}

	for _, cinema := range <-chanList {
		result.Data = append(result.Data, &manager_api.CinemaRoomModel{
			Uuid:           cinema.Uuid,
			MovieTheaterId: cinema.MovieTheaterId,
			Code:           cinema.Code,
			CreatedAt:      cinema.CreatedAt.String(),
			UpdatedAt:      cinema.UpdatedAt.String(),
		})
	}

	return result, nil
}
