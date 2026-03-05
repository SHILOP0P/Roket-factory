package v1

import (
	"order/internal/service"
)

type api struct {

	OrderService	service.OrderService
}

func NewAPI(orderService service.OrderService) *api{
	return &api{
		OrderService: orderService,
	}
}