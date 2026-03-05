package v1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"payment/internal/converter"

	payment "shared/pkg/proto/payment/v1"
)

func (s *api) PayOrder(ctx context.Context, req *payment.PayOrderRequest) (*payment.PayOrderResponse, error) {
	// Можно использовать эти поля для доп. логики/валидации
	orderUUID := req.GetOrderUuid()
	userUUID := req.GetUserUuid()
	paymentMethod := req.GetPaymentMethod()

	transactionUUID, err := s.PaymentService.PayOrder(ctx, orderUUID, userUUID, converter.PaymentMethodProtoToModel(paymentMethod))
	if err!=nil{
		return nil, status.Errorf(codes.Internal, "Failed transaction")
	}

	// Генерация UUID транзакции (v4)


	return &payment.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil
}