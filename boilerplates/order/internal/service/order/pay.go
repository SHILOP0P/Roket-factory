package order

import(
	"context"
	"order/internal/model"
)

func (s *service) PayOrder(ctx context.Context, orderUpdate model.UpdateOrder)(error){
	err := s.repository.UpdateOrder(ctx, orderUpdate)
	if err!=nil{
		return err
	}
	return nil
}