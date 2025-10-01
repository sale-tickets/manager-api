package health_controller

import (
	"context"

	manager_api "github.com/duyhung2k4/sale-tickets-golang-common/manager-api/proto"
)

func (m *managerService) Ping(ctx context.Context, req *manager_api.Request) (*manager_api.Response, error) {
	return &manager_api.Response{
		Mess: "oke",
	}, nil
}
