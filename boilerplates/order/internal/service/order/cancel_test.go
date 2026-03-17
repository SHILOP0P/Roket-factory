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

	s.orderRepository.
		On("GetOrderByUUID", s.ctx, orderUUID).
		Return(model.Order{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
			Status: model.OrderStatusPENDINGPAYMENT,
		}, nil).
		Once()

	s.orderRepository.On("UpdateOrder", s.ctx, cancelInfo).Return(nil)

	err:= s.service.CancelOrder(s.ctx, cancelInfo)
	s.NoError(err)
}

func (s *ServiceSuit) TestCancelOrderError(){
	var(
		getErr = gofakeit.Error()
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		cancelInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
		}
	)
	s.orderRepository.
		On("GetOrderByUUID", s.ctx, orderUUID).
		Return(model.Order{}, getErr).
		Once()
	
	err := s.service.CancelOrder(s.ctx, cancelInfo)
	s.Error(err)
	s.ErrorIs(err, getErr)
}
