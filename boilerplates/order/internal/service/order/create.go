package order

import(
	"context"
	"order/internal/model"
)

func (s *service) CreateOrder(ctx context.Context, order model.Order)(model.Order, error){
	repoorder, err := s.repository.CreateOrder(ctx, order)
	if err!=nil{
		return model.Order{}, err
	}
	return repoorder, nil
}