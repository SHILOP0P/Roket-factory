package v1

import (
	"context"
	"errors"
	"order/internal/converter"
	"order/internal/model"

	orderv1 "shared/pkg/openapi/order/v1"
)

func (a *api) GetOrderByUUID(ctx context.Context, params orderv1.GetOrderByUUIDParams) (orderv1.GetOrderByUUIDRes, error) {
	
	order, err := a.OrderService.GetOrder(ctx, converter.GetOrderByUUIDParamsToString(params))
	if err!=nil{
		if errors.Is(err, model.ErrOrderNotFound) {
			return &orderv1.GetOrderByUUIDNotFound{
				Code:    404,
				Message: "order not found",
			}, nil
		}

		return &orderv1.GetOrderByUUIDInternalServerError{
			Code:    500,
			Message: "internal error",
		}, nil
	}
	
	res, err := converter.OrderToGetOrderByUUIDRes(order)
	if err != nil {
		return &orderv1.GetOrderByUUIDInternalServerError{
			Code:    500,
			Message: "failed to build response",
		}, nil
	}
	return res, nil
}