package order

import (
	"context"
	"database/sql"
	"errors"
	"order/internal/model"
	"order/internal/repository/converter"
	repoModel "order/internal/repository/model"
	"strconv"
	"strings"

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
	setParts := make([]string, 0)
	args := make([]any, 0)
	idx := 1
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

	if repoOrderUpdate.PartUUIDs != nil {
		setParts = append(setParts, "part_uuids = $"+strconv.Itoa(idx))
		args = append(args, pq.Array(*repoOrderUpdate.PartUUIDs))
		idx++
	}
	if repoOrderUpdate.PaymentMethod != nil {
		setParts = append(setParts, "payment_method = $"+strconv.Itoa(idx))
		args = append(args, *repoOrderUpdate.PaymentMethod)
		idx++
	}
	if repoOrderUpdate.Status != nil {
		setParts = append(setParts, "order_status = $"+strconv.Itoa(idx))
		args = append(args, *repoOrderUpdate.Status)
		idx++
	}
	if repoOrderUpdate.TransactionUUID != nil {
		setParts = append(setParts, "transaction_uuid = $"+strconv.Itoa(idx))
		args = append(args, *repoOrderUpdate.TransactionUUID)
		idx++
	}
	if repoOrderUpdate.TotalPrice != nil {
		setParts = append(setParts, "total_price = $"+strconv.Itoa(idx))
		args = append(args, *repoOrderUpdate.TotalPrice)
		idx++
	}

	updateQuery := `
		UPDATE orders
		SET ` + strings.Join(setParts, ", ") + `
		WHERE order_uuid = $` + strconv.Itoa(idx)
	args = append(args, orderUpdate.OrderUUID)

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
