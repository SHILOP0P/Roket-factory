package order

import(
	"github.com/brianvoe/gofakeit/v6"

	"order/internal/model"
)

func(s *ServiceSuit) TestCancelOrderSuccess(){
	var(
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		status = model.OrderStatusCANCELLED

		cancelInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
			Status: &status,
		}
	)

	s.orderRepository.On("UpdateOrder", s.ctx, cancelInfo).Return(nil)

	err:= s.service.CancelOrder(s.ctx, cancelInfo)
	s.NoError(err)
}

func (s *ServiceSuit) TestCancelOrderError(){
	var(
		repoErr = gofakeit.Error()
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		cancelInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
		}
	)
	s.orderRepository.On("UpdateOrder", s.ctx, cancelInfo).Return(repoErr)
	
	err := s.service.CancelOrder(s.ctx, cancelInfo)
	s.Error(err)
	s.ErrorIs(err, repoErr)
}
