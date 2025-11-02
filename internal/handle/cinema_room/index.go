package cinemaroom_controller

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	cinemaroom_service "github.com/sale-tickets/manager-api/internal/service/cinema_room"
	theaterseating_service "github.com/sale-tickets/manager-api/internal/service/theater_seating"
)

type cinemaRoomController struct {
	service               cinemaroom_service.CinemaRoomService
	theaterSeatingService theaterseating_service.TheaterSeatingService
	manager_api.UnimplementedCinemaRoomServiceServer
}

func NewHandle(service cinemaroom_service.CinemaRoomService, theaterSeatingService theaterseating_service.TheaterSeatingService) manager_api.CinemaRoomServiceServer {
	return &cinemaRoomController{
		service:               service,
		theaterSeatingService: theaterSeatingService,
	}
}
