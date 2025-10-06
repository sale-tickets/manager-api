package view

import manager_api "github.com/sale-tickets/golang-common/manager-api/proto"

// api: create theater seating
type CreateTheaterSeatingReq struct {
	*manager_api.CreateTheaterSeatingReq
}

func (v *CreateTheaterSeatingReq) Validate() error {
	return nil
}

// api: update theater seating
type UpdateTheaterSeatingReq struct {
	*manager_api.UpdateTheaterSeatingReq
}

func (v *UpdateTheaterSeatingReq) Validate() error {
	return nil
}

// api: list theater seating
type ListTheaterSeatingReq struct {
	*manager_api.ListTheaterSeatingReq
}

func (v *ListTheaterSeatingReq) Validate() error {
	return nil
}
