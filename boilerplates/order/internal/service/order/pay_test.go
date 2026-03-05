package order

import(
	"github.com/brianvoe/gofakeit/v6"

	"order/internal/model"
)

func(s *ServiceSuit) TestPayOrderSuccess(){
	var(
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		status = model.OrderStatusPAID

		payInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
			Status: &status,
		}
	)

	s.orderRepository.On("UpdateOrder", s.ctx, payInfo).Return(nil)

	err:= s.service.PayOrder(s.ctx, payInfo)
	s.NoError(err)
}

func (s *ServiceSuit) TestPayOrderError(){
	var(
		repoErr = gofakeit.Error()
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		payInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
		}
	)
	s.orderRepository.On("UpdateOrder", s.ctx, payInfo).Return(repoErr)
	
	err := s.service.PayOrder(s.ctx, payInfo)
	s.Error(err)
	s.ErrorIs(err, repoErr)
}
