package order

import (
	"context"
	"order/internal/model"
	"order/internal/repository/converter"
)

func (r *repository) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	r.mu.RLock()
	defer r.mu.Unlock()

	repoOrder := converter.OrderToRepoModel(order)
	r.storage[repoOrder.OrderUUID] = &repoOrder

	repoOrderNew, ok := r.storage[repoOrder.OrderUUID]
	if !ok{
		return model.Order{}, model.ErrFailCreated
	}

	modelOrder := converter.OrderToModel(*repoOrderNew)
	
	return modelOrder, nil
}