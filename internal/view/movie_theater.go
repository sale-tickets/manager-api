package view

import (
	"errors"

	manager_api "github.com/duyhung2k4/sale-tickets-golang-common/manager-api/proto"
)

type CreateMovieTheaterReq struct {
	*manager_api.CreateMovieTheaterReq
	CreaterId string
}

func (v *CreateMovieTheaterReq) Validate() error {
	if v.Name == "" {
		return errors.New("invalid name")
	}

	if v.Address == "" {
		return errors.New("address name")
	}

	return nil
}

type GetListMovieTheaterReq struct {
	*manager_api.GetListMovieTheaterReq
}

func (v *GetListMovieTheaterReq) Validate() error {
	if v == nil {
		return errors.New("request data nil")
	}
	if v.FilterBase == nil {
		return errors.New("filter_base data nil")
	}
	if v.Filter == nil {
		return errors.New("filter data nil")
	}
	if v.FilterBase.Limit < 0 {
		return errors.New("invalid limit")
	}
	if v.FilterBase.Offset < 0 {
		return errors.New("invalid offset")
	}

	return nil
}

type DetailMovieTheaterReq struct {
	*manager_api.DetailMovieTheaterReq
}

type UpdateMovieTheaterReq struct {
	*manager_api.UpdateMovieTheaterReq
}

func (v *UpdateMovieTheaterReq) Validate() error {
	if v == nil {
		return errors.New("request data nil")
	}
	if v.Name == "" {
		return errors.New("name invalid")
	}
	if v.Address == "" {
		return errors.New("address invalid")
	}
	return nil
}
