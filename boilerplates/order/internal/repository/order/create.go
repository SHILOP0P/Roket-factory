package order

import (
	"context"
	"order/internal/model"
	"order/internal/repository/converter"
	repoModel "order/internal/repository/model"

	"github.com/lib/pq"
)

func (r *repository) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	repoOrder := converter.OrderToRepoModel(order)
	var repoOrderNew repoModel.Order
	createQuery := `
		INSERT INTO orders (order_uuid, user_uuid, part_uuids, total_price, transaction_uuid, payment_method, order_status)
		VALUES ($1, $2, $3, $4, $5, $6, $7) returning order_uuid, user_uuid, part_uuids, total_price, transaction_uuid, payment_method, order_status
	`
	err := r.db.QueryRowContext(ctx, createQuery,
		repoOrder.OrderUUID,
		repoOrder.UserUUID,
		pq.Array(repoOrder.PartUUIDs),
		repoOrder.TotalPrice,
		repoOrder.TransactionUUID,
		repoOrder.PaymentMethod,
		repoOrder.Status,
	).Scan(
		&repoOrderNew.OrderUUID,
		&repoOrderNew.UserUUID,
		pq.Array(&repoOrderNew.PartUUIDs),
		&repoOrderNew.TotalPrice,
		&repoOrderNew.TransactionUUID,
		&repoOrderNew.PaymentMethod,
		&repoOrderNew.Status,
	)
	if err != nil {
		return model.Order{}, err
	}
	modelOrder := converter.OrderToModel(repoOrderNew)
	
	return modelOrder, nil
}