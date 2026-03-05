package order

import (
	"context"
	"order/internal/model"
	"order/internal/repository/converter"
)

func (r *repository) GetOrderByUUID(_ context.Context, orderUUID string) (model.Order, error) {
	r.mu.RLock()
	o, ok := r.storage[orderUUID]
	r.mu.RUnlock()
	if !ok {
		return model.Order{}, model.ErrOrderNotFound
	}

	return converter.OrderToModel(*o), nil
}