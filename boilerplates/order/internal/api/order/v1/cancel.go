package v1

import (
	"context"
	"errors"
	"order/internal/model"
	"order/internal/converter"
	orderV1 "shared/pkg/openapi/order/v1"
)

func (a *api) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	
	err :=a.OrderService.CancelOrder(ctx, converter.CancelOrderParamsToModel(params))
	if err!=nil{
		if errors.Is(err, model.ErrOrderNotFound) {
				return &orderV1.CancelOrderNotFound{Code: 404, Message: "Order not found"}, nil
		}
		return &orderV1.CancelOrderInternalServerError{Code: 500, Message: "internal error"}, nil
	}
	

	return &orderV1.CancelOrderNoContent{}, nil
}