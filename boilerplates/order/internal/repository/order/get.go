package order

import (
	"context"
	"database/sql"
	"errors"
	"order/internal/model"
	"order/internal/repository/converter"
	repoModel "order/internal/repository/model"

	"github.com/lib/pq"
)

func (r *repository) GetOrderByUUID(ctx context.Context, orderUUID string) (model.Order, error) {
	var repoOrder repoModel.Order
	getQuery := `
		SELECT order_uuid, user_uuid, part_uuids, total_price, transaction_uuid, payment_method, order_status
		FROM orders
		WHERE order_uuid = $1
	`
	err := r.db.QueryRowContext(ctx, getQuery, orderUUID).Scan(
		&repoOrder.OrderUUID,
		&repoOrder.UserUUID,
		pq.Array(&repoOrder.PartUUIDs),
		&repoOrder.TotalPrice,
		&repoOrder.TransactionUUID,
		&repoOrder.PaymentMethod,
		&repoOrder.Status,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Order{}, model.ErrOrderNotFound
		}
		return model.Order{}, err
	}

	return converter.OrderToModel(repoOrder), nil
}
