package cinemaroom_controller

import manager_api "github.com/sale-tickets/golang-common/manager-api/proto"

type cinemaRoomController struct {
	manager_api.UnimplementedCinemaRoomServiceServer
}

func NewHandle() manager_api.CinemaRoomServiceServer {
	return &cinemaRoomController{}
}
