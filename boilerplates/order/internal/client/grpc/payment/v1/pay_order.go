package v1

import (
	"context"
	"order/internal/client/converter"
	"order/internal/model"
	payment_v1 "shared/pkg/proto/payment/v1"
)

func (p *paymentClient) PayOrder(ctx context.Context, orderUUID string, userUUID string, method model.PaymentMethod) (string, error){
	req := &payment_v1.PayOrderRequest{
		OrderUuid: orderUUID,
		UserUuid: userUUID,
		PaymentMethod: converter.PaymentMethodModelToProto(method),
	}
	res, err := p.client.PayOrder(ctx, req)
	if err!=nil{
		return "", err
	}
	return res.TransactionUuid, nil
}