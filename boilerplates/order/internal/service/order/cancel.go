package order

import(
	"context"
	"order/internal/model"
)

func (s *service) CancelOrder(ctx context.Context, orderUpdate model.UpdateOrder)(error){

	order, err := s.repository.GetOrderByUUID(ctx, orderUpdate.OrderUUID)
	if err != nil{
		return err
	}
	if order.Status==model.OrderStatusCANCELLED{
		return nil
	}
	if order.Status==model.OrderStatusPAID{
		return model.ErrFailCancel
	}
	err = s.repository.UpdateOrder(ctx, orderUpdate)
	if err!=nil{
		return err
	}
	return nil
}