package movietheater_service

import (
	"github.com/sale-tickets/manager-api/internal/view"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

func (s *movieTheaterService) GetList(req view.GetListMovieTheaterReq) (*manager_api.GetListMovieTheaterRes, error) {
	var chanErrCount, chanErrGetList = make(chan error, 1), make(chan error, 1)
	var chanCount = make(chan int32, 1)
	var chanList = make(chan []*manager_api.MovieTheaterRes, 1)

	go func() {
		count, err := s.movieTheaterRepo.CountList(req)
		chanErrCount <- err
		chanCount <- count
	}()

	go func() {
		movieTheaters, err := s.movieTheaterRepo.GetList(req)

		result := []*manager_api.MovieTheaterRes{}
		for _, item := range movieTheaters {
			result = append(result, &manager_api.MovieTheaterRes{
				Uuid:      item.Uuid,
				Name:      item.Name,
				Address:   item.Address,
				CreaterId: item.CreaterId,
			})
		}

		chanErrGetList <- err
		chanList <- result
	}()

	errCount := <-chanErrCount
	if errCount != nil {
		return nil, errCount
	}
	errList := <-chanErrGetList
	if errList != nil {
		return nil, errList
	}

	count := <-chanCount
	list := <-chanList
	result := &manager_api.GetListMovieTheaterRes{
		Total: count,
		Data:  list,
	}

	return result, nil
}
