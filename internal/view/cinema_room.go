package view

import (
	"errors"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

type CreateCinemaRoomReq struct {
	CreaterId string
	*manager_api.CreateCinemaRoomReq
}

func (v *CreateCinemaRoomReq) Validate() error {
	if v == nil {
		return errors.New("request data nil")
	}
	if v.MovieTheaterId == "" {
		return errors.New("movie_theater_id invalid")
	}
	if v.Code == "" {
		return errors.New("code invalid")
	}
	return nil
}

type UpdateCinemaRoomReq struct {
	CreaterId string
	*manager_api.UpdateCinemaRoomReq
}

func (v *UpdateCinemaRoomReq) Validate() error {
	if v == nil {
		return errors.New("request data nil")
	}
	if v.Uuid == "" {
		return errors.New("uuid invalid")
	}
	if v.Code == "" {
		return errors.New("code invalid")
	}
	return nil
}

type ListCinemaRoomReq struct {
	CreateId string
	*manager_api.ListCinemaRoomReq
}

func (v *ListCinemaRoomReq) Validate() error {
	if v == nil {
		return errors.New("request data nil")
	}
	if v.FilterBase == nil {
		return errors.New("filter_base invalid")
	}
	if v.Filter == nil {
		return errors.New("filter invalid")
	}
	return nil
}

type DetailCinemaRoomReq struct {
	CreaterId string
	*manager_api.DetailCinemaRoomReq
}

func (v *DetailCinemaRoomReq) Validate() error {
	return nil
}
