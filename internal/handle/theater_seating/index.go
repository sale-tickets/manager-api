package theaterseating_controller

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	theaterseating_service "github.com/sale-tickets/manager-api/internal/service/theater_seating"
)

type theaterSeatingController struct {
	service theaterseating_service.TheaterSeatingService
	manager_api.UnimplementedTheaterSeatingServer
}

func NewHandle() manager_api.TheaterSeatingServer {
	return &theaterSeatingController{
		service: theaterseating_service.NewTheaterSeatingService(),
	}
}
