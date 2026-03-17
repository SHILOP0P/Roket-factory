package order

import (
	"order/internal/repository"
	service_interface "order/internal/service"
)

type service struct{
	repository repository.OrderRepository
	inventoryClient service_interface.InventoryClient
	paymentClient service_interface.PaymentClient
}

func NewOrderService(repository repository.OrderRepository, inventoryClient service_interface.InventoryClient, paymentClient service_interface.PaymentClient) *service {
	return &service{
		repository: repository,
		inventoryClient: inventoryClient,
		paymentClient: paymentClient,
	}
}