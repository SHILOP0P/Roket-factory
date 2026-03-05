package order

import (
	repoModel"order/internal/repository/model"
	"sync"
)

type repository struct{
	mu sync.RWMutex
	storage map[string] *repoModel.Order
}

func NewOrderRepository() *repository {
	return &repository{
		storage: make(map[string]*repoModel.Order),
	}
}