package order

import (
	"context"
	"order/internal/model"
	"order/internal/repository/converter"
)

func (r *repository) UpdateOrder(ctx context.Context, orderUpdate model.UpdateOrder) (error){
	r.mu.Lock()
	defer r.mu.Unlock()
	repoOrderUpdate := converter.OrderUpdateToRepoModel(orderUpdate)
	order, ok := r.storage[repoOrderUpdate.OrderUUID]
	if !ok {
		return model.ErrOrderNotFound
	}
	if repoOrderUpdate.PartUUIDs!=nil && repoOrderUpdate.PaymentMethod!=nil && repoOrderUpdate.Status!=nil{
		return nil
	}
	if repoOrderUpdate.PartUUIDs!=nil{
		order.PartUUIDs = *orderUpdate.PartUUIDs
	}
	if repoOrderUpdate.PaymentMethod!=nil{
		order.PaymentMethod = repoOrderUpdate.PaymentMethod
	}
	if repoOrderUpdate.Status!=nil{
		order.Status = *repoOrderUpdate.Status
	}
	r.storage[order.OrderUUID] = order
	return nil
}