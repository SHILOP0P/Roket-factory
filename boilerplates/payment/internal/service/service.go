package service

import (
	"context"
	"payment/internal/model"
)

type PaymentService interface {
	PayOrder(ctx context.Context, orderUUID string, userUUID string, method model.PaymentMethod) (string, error)
}
