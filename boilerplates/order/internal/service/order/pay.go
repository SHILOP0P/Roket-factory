package order

import (
	"context"
	"log"
	"order/internal/model"
)

func (s *service) PayOrder(ctx context.Context, orderUpdate model.UpdateOrder)(error){
	if orderUpdate.PaymentMethod==nil{
		log.Printf("Payment method in PayRequest is nil")
		return model.ErrFailPayed
	}
	
	order, err := s.repository.GetOrderByUUID(ctx, orderUpdate.OrderUUID)
	if err!=nil{
		return err
	}
	orderUpdate.UserUUID = order.UserUUID

	transactionUUID, err := s.paymentClient.PayOrder(ctx, orderUpdate.OrderUUID, orderUpdate.UserUUID, *orderUpdate.PaymentMethod)
	if err!=nil{
		log.Printf("transaction failed: %v", err)
		return err
	}
	status := model.OrderStatusPAID
	orderUpdate.TransactionUUID = &transactionUUID
	orderUpdate.Status = &status
	err = s.repository.UpdateOrder(ctx, orderUpdate)
	if err!=nil{
		log.Printf("failed pay order: %v", err)
		return err
	}
	return nil
}
