package view

import (
	"errors"

	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

type CreateMovieReq struct {
	*manager_api.CreateMovieReq
}

func (v *CreateMovieReq) Validate() error {
	if v.Data == nil {
		return errors.New("data not found")
	}
	return nil
}

type UpdateMovieReq struct {
	*manager_api.UpdateMovieReq
}

func (v *UpdateMovieReq) Validate() error {
	if v.Data == nil {
		return errors.New("data not found")
	}
	return nil
}

type ListMovieReq struct {
	*manager_api.ListMovieReq
}
