package v1

import (
	"context"
	"errors"
	"order/internal/converter"
	"order/internal/model"

	orderv1 "shared/pkg/openapi/order/v1"
)

func (a *api) PayOrder(ctx context.Context, req *orderv1.PayOrderRequest, params orderv1.PayOrderParams) (orderv1.PayOrderRes, error) {
	update := converter.PayOrderRequestToModel(params, req)

	if err := a.OrderService.PayOrder(ctx, update); err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &orderv1.PayOrderNotFound{Code: 404, Message: "order not found"}, nil
		}
		return &orderv1.PayOrderInternalServerError{Code: 500, Message: "internal error"}, nil
	}

	order, err := a.OrderService.GetOrder(ctx, update.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &orderv1.PayOrderNotFound{Code: 404, Message: "order not found"}, nil
		}
		return &orderv1.PayOrderInternalServerError{Code: 500, Message: "internal error"}, nil
	}

	res, err := converter.OrderToPayOrderResponse(order)
	if err != nil {
		return &orderv1.PayOrderInternalServerError{Code: 500, Message: "failed to build response"}, nil
	}

	return res, nil
}
