package v1

import (
	paymentV1 "shared/pkg/proto/payment/v1"
	"payment/internal/service"
)

type api struct {
	paymentV1.UnimplementedPaymentServiceServer

	PaymentService	service.PaymentService
}

func NewAPI(paymentService service.PaymentService) *api{
	return &api{
		PaymentService: paymentService,
	}
}