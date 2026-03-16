package service

import(
	"context"
	"order/internal/model"
)

type OrderService interface{
	CreateOrder(ctx context.Context, order model.Order)(model.Order, error)
	GetOrder(ctx context.Context, orderUUID string)(model.Order, error)
	PayOrder(ctx context.Context, orderUpdate model.UpdateOrder)(error)
	CancelOrder(ctx context.Context, orderUpdate model.UpdateOrder)(error)
}

type InventoryClient interface{
	GetInventoryModels(ctx context.Context, uuids []string)([]model.Part, error)
}

type PaymentClient interface{
	PayOrder(ctx context.Context, orderUUID string, userUUID string, method model.PaymentMethod) (string, error)
}
