package order

import(
	"github.com/brianvoe/gofakeit/v6"
	"order/internal/model"
)

func(s *ServiceSuit) TestCreateOrderSuccess(){
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
	s.orderRepository.On("CreateOrder", s.ctx, order).Return(order, nil)

	res, err := s.service.CreateOrder(s.ctx, order)
	s.Require().NoError(err)
	s.Require().Equal(order, res)
}

func (s *ServiceSuit) TestCreateOrderError(){
	var(
		repoError = gofakeit.Error()
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
	s.orderRepository.On("CreateOrder", s.ctx, order).Return(model.Order{}, repoError)

	res, err := s.service.CreateOrder(s.ctx, order)
	s.Error(err)
	s.ErrorIs(err, repoError)
	s.Empty(res)
}




func fakePartsUUIDs(n int)[]string{
	listPartsUUIDs := make([]string, n)

	for i := 0; i < n; i++ {
		partUUID := gofakeit.UUID()
		listPartsUUIDs[i] = partUUID
	}
	return listPartsUUIDs
}