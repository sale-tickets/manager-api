package health_controller

import (
	manager_api "github.com/sale-tickets/golang-common/manager-api/proto"
)

type managerService struct {
	manager_api.UnimplementedHealthServer
}

func NewHandle() manager_api.HealthServer {
	return &managerService{}
}
