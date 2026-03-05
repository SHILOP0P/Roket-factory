package order

import(
	"github.com/brianvoe/gofakeit/v6"
	"order/internal/model"
)

func(s *ServiceSuit) TestGetOrderSuccess(){
	var(
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		partUUIDs = fakePartsUUIDs(gofakeit.Number(1, 5))
		totalPrice = gofakeit.Float64()
		transactionUUID = gofakeit.UUID()
		paymentMethod = model.PaymentMethod(gofakeit.Number(0, 4))
		orderStatuses = []string{
			string(model.OrderStatusPENDINGPAYMENT),
			string(model.OrderStatusPAID),
			string(model.OrderStatusCANCELLED),
		}
		status = model.OrderStatus(gofakeit.RandomString(orderStatuses))

		order = model.Order{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
			PartUUIDs: partUUIDs,
			TotalPrice: totalPrice,
			TransactionUUID: &transactionUUID,
			PaymentMethod: &paymentMethod,
			Status: status,
		}
	)
	s.orderRepository.On("GetOrderByUUID", s.ctx, orderUUID).Return(order, nil)

	res, err := s.service.GetOrder(s.ctx, orderUUID)
	s.Require().NoError(err)
	s.Require().Equal(order, res)
}

func (s *ServiceSuit) TestGetOrderError(){
	var(
		repoError = gofakeit.Error()
		orderUUID = gofakeit.UUID()
	)
	s.orderRepository.On("GetOrderByUUID", s.ctx, orderUUID).Return(model.Order{}, repoError)

	res, err := s.service.GetOrder(s.ctx, orderUUID)
	s.Error(err)
	s.ErrorIs(err, repoError)
	s.Empty(res)
}

