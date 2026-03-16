package v1

import (
	"context"
	"errors"
	"order/internal/converter"
	"order/internal/model"

	orderv1 "shared/pkg/openapi/order/v1"
)

func (a *api) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (orderv1.CreateOrderRes, error) {
	in := converter.CreateOrderRequestToModel(req)

	createdOrder, err := a.OrderService.CreateOrder(ctx, in)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrFailCreated):
			return &orderv1.CreateOrderInternalServerError{Code: 500, Message: err.Error()}, nil
		default:
			return &orderv1.CreateOrderNotFound{Code: 404, Message: err.Error()}, nil
		}
	}

	res, err := converter.OrderToCreateOrderResponse(createdOrder)
	if err != nil {
		return &orderv1.CreateOrderInternalServerError{Code: 500, Message: err.Error()}, nil
	}

	return res, nil
}
