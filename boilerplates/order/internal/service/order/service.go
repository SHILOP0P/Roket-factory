package order

import (
	"order/internal/repository"
)

type service struct{
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) *service {
	return &service{
		repository: repository,
	}
}