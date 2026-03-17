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

func (r *repository) UpdateOrder(ctx context.Context, orderUpdate model.UpdateOrder) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			return
		}
	}()

	repoOrderUpdate := converter.OrderUpdateToRepoModel(orderUpdate)
	var repoOrder repoModel.Order
	args := make([]any, 0)
	getQuery := `
		SELECT order_uuid, user_uuid, part_uuids, total_price, transaction_uuid, payment_method, order_status
		FROM orders
		WHERE order_uuid = $1
		for update
	`
	err = tx.QueryRowContext(ctx, getQuery, orderUpdate.OrderUUID).Scan(
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
			return model.ErrOrderNotFound
		}
		return err
	}

	if repoOrderUpdate.PartUUIDs == nil && repoOrderUpdate.PaymentMethod == nil && repoOrderUpdate.Status == nil && repoOrderUpdate.TotalPrice == nil && repoOrderUpdate.TransactionUUID == nil {
		return nil
	}

	var partUUIDs any
	if repoOrderUpdate.PartUUIDs != nil {
		partUUIDs = pq.Array(*repoOrderUpdate.PartUUIDs)
	}
	updateQuery := `
		UPDATE orders
		SET part_uuids = COALESCE($1::uuid[], part_uuids),
			payment_method = COALESCE($2::integer, payment_method),
			order_status = COALESCE($3::varchar(20), order_status),
			transaction_uuid = COALESCE($4::uuid, transaction_uuid),
			total_price = COALESCE($5::numeric(10, 2), total_price)
		WHERE order_uuid = $6
	`
	args = append(args,
		partUUIDs,
		int32Value(repoOrderUpdate.PaymentMethod),
		stringValue(repoOrderUpdate.Status),
		stringValue(repoOrderUpdate.TransactionUUID),
		float64Value(repoOrderUpdate.TotalPrice),
		orderUpdate.OrderUUID,
	)

	result, err := tx.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return model.ErrOrderNotFound
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func int32Value[T ~int32](value *T) any {
	if value == nil {
		return nil
	}
	return int32(*value)
}

func stringValue[T ~string](value *T) any {
	if value == nil {
		return nil
	}
	return string(*value)
}

func float64Value(value *float64) any {
	if value == nil {
		return nil
	}
	return *value
}
