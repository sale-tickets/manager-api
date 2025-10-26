package theaterseating_service

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	"github.com/sale-tickets/golang-common/model"

	"github.com/sale-tickets/manager-api/internal/view"
)

func (s *theaterSeatingService) List(req *view.ListTheaterSeatingReq) (*manager_api.ListTheaterSeatingRes, error) {
	var countChan, listChan = make(chan int64, 1), make(chan []*model.TheaterSeating, 1)
	var errCountChan, errListChan = make(chan error, 1), make(chan error, 1)

	go func() {
		count, err := s.repo.Count(req)
		countChan <- count
		errCountChan <- err
	}()

	go func() {
		list, err := s.repo.List(req)
		listChan <- list
		errListChan <- err
	}()

	if err := <-errCountChan; err != nil {
		return nil, err
	}
	if err := <-errListChan; err != nil {
		return nil, err
	}

	result := &manager_api.ListTheaterSeatingRes{
		Total: 0,
		Data:  make([]*manager_api.TheaterSeatingModel, 0),
	}

	count := <-countChan
	if count == 0 {
		return result, nil
	}
	result.Total = int32(count)

	for _, item := range <-listChan {
		result.Data = append(result.Data, &manager_api.TheaterSeatingModel{
			Uuid:         item.Uuid,
			CinemaRoomId: item.CinemaRoomId,
			Name:         item.Name,
			RowIndex:     item.RowIndex,
			ColIndex:     item.ColIndex,
		})
	}

	return result, nil
}
