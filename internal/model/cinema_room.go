package model

type CinemaRoom struct {
	Base
	MovieTheaterId string `json:"movie_theater_id" gorm:"uniqueIndex:idx_movie_theater_id_code"`
	Code           string `json:"code" gorm:"uniqueIndex:idx_movie_theater_id_code"`
}
