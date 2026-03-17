package order

import (
	"context"
	"log"
	"order/internal/model"

	"github.com/google/uuid"
)

func (s *service) CreateOrder(ctx context.Context, order model.Order)(model.Order, error){
	order.OrderUUID = uuid.NewString()
	order.Status = model.OrderStatusPENDINGPAYMENT
	parts, err := s.inventoryClient.GetInventoryModels(ctx, order.PartUUIDs)
	if err!= nil{
		log.Printf("Getting parts failed: %v", err)
		return model.Order{}, err
	}
	if len(parts)!=len(order.PartUUIDs){
		log.Printf("Count TotalPrice failed: parts count mismatch")
		return model.Order{}, model.ErrFailCreated
	}
	for i := 0; i < len(parts); i++ {
		if parts[i].Uuid!=order.PartUUIDs[i]{
			return model.Order{}, nil
		}
	}
	order.TotalPrice = 0
	for _, part := range parts{
		order.TotalPrice+=part.Price
	}
	repoorder, err := s.repository.CreateOrder(ctx, order)
	if err!=nil{
		return model.Order{}, err
	}
	return repoorder, nil
}
