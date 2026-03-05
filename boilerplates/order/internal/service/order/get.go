package order

import(
	"context"
	"order/internal/model"
)

func (s *service) GetOrder(ctx context.Context, orderUUID string)(model.Order, error){
	order, err := s.repository.GetOrderByUUID(ctx, orderUUID)
	if err!=nil{
		return model.Order{}, err
	}
	return order, nil
}