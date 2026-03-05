package v1

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"order/internal/model"
	"order/internal/converter"
	orderv1 "shared/pkg/openapi/order/v1"
)

func (a *api) CancelOrder(ctx context.Context, params orderv1.CancelOrderParams) (orderv1.CancelOrderRes, error) {
	
	err :=a.OrderService.CancelOrder(ctx, converter.CancelOrderParamsToModel(params))
	if err!=nil{
		if errors.Is(err, model.ErrOrderNotFound) {
				return nil, status.Errorf(codes.NotFound, "Order with UUID %s not found", params.OrderUUID)
			}
		// return &orderv1.PayOrderInternalServerError{Code: 500, Message: "internal error"}, nil
		return  nil, status.Errorf(codes.Internal, "internal error")
	}
	

	return &orderv1.CancelOrderNoContent{}, nil
}