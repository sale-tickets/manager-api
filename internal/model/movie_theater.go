package model

type MovieTheater struct {
	Base
	CreaterId string `json:"creater_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}
