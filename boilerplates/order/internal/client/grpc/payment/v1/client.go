package v1

import(
	payment_v1 "shared/pkg/proto/payment/v1"
	"order/internal/service"
)

type paymentClient struct{
	client payment_v1.PaymentServiceClient
}

func NewPaymentClient(client payment_v1.PaymentServiceClient) service.PaymentClient{
	return &paymentClient{
		client: client,
	}
}