package payment

import(
	"context"
	"payment/internal/model"
	"github.com/google/uuid"
	"log"
)

func (s *service) PayOrder(_ context.Context, orderUUID, userUUID string, payMethod model.PaymentMethod) (string, error) {
	// Генерация UUID транзакции (v4)
	transactionUUID := uuid.NewString()

	log.Printf(
		"Оплата прошла успешно, transaction_uuid: %s, order_uuid: %s, user_uuid: %s, payment_method: %v",
		transactionUUID,
		orderUUID,
		userUUID,
		payMethod,
	)

	return transactionUUID, nil
}