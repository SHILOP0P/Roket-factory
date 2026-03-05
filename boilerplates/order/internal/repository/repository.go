package repository

import(
	"context"
	"order/internal/model"
)

type OrderRepository interface{
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderByUUID(ctx context.Context, orderUUID string) (model.Order, error)
	UpdateOrder(ctx context.Context, orderUpdate model.UpdateOrder) (error)
}