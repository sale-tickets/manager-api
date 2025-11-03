package showtime_controller

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
	showtime_repo "github.com/sale-tickets/manager-api/internal/repo/showtime"
)

type ShowtimeController struct {
	showtimeRepo *showtime_repo.ShowtimeRepo
	manager_api.UnimplementedShowtimeServer
}

func NewHandle(showtimeRepo *showtime_repo.ShowtimeRepo) manager_api.ShowtimeServer {
	return &ShowtimeController{
		showtimeRepo: showtimeRepo,
	}
}
