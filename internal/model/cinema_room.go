package model

type CinemaRoom struct {
	Base
	MovieTheaterId string `json:"movie_theater_id"`
	Code           string `json:"code" gorm:"uniqueIndex:idx_cinema_rooms_uuid"`
}
