package model

import "github.com/lib/pq"

type TheaterSeating struct {
	Base
	CinemaRoomId   string         `json:"cinema_room_id"`
	GroupSeatingId pq.StringArray `json:"group_seating_id" gorm:"type:text[]"`
	Name           string         `json:"name"`
	RowIndex       string         `json:"row_index"`
	ColIndex       string         `json:"col_index"`
}
