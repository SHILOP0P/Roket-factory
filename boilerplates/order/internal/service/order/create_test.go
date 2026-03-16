package order

import(
	"context"
	"order/internal/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func(s *ServiceSuit) TestCreateOrderSuccess(){
	var(
		userUUID = gofakeit.UUID()
		partUUIDs = fakePartsUUIDs(gofakeit.Number(1, 5))
		transactionUUID = gofakeit.UUID()
		paymentMethod = model.PaymentMethod(gofakeit.Number(0, 4))
		orderStatuses = []string{
			string(model.OrderStatusPENDINGPAYMENT),
			string(model.OrderStatusPAID),
			string(model.OrderStatusCANCELLED),
		}
		status = model.OrderStatus(gofakeit.RandomString(orderStatuses))

		order = model.Order{
			UserUUID: userUUID,
			PartUUIDs: partUUIDs,
			TransactionUUID: &transactionUUID,
			PaymentMethod: &paymentMethod,
			Status: status,
		}
		parts = []model.Part{
			{Uuid: partUUIDs[0], Price: 10.5},
		}
		expectedTotalPrice = 10.5
	)
	for _, partUUID := range partUUIDs[1:] {
		price := gofakeit.Float64Range(1, 100)
		parts = append(parts, model.Part{Uuid: partUUID, Price: price})
		expectedTotalPrice += price
	}

	s.inventoryClient.
		On("GetInventoryModels", s.ctx, partUUIDs).
		Return(parts, nil).
		Once()

	s.orderRepository.
		On("CreateOrder", s.ctx, mock.MatchedBy(func(in model.Order) bool {
			_, err := uuid.Parse(in.OrderUUID)
			return err == nil &&
				in.UserUUID == order.UserUUID &&
				len(in.PartUUIDs) == len(order.PartUUIDs) &&
				in.TotalPrice == expectedTotalPrice &&
				*in.TransactionUUID == *order.TransactionUUID &&
				*in.PaymentMethod == *order.PaymentMethod &&
				in.Status == model.OrderStatusPENDINGPAYMENT
		})).
		Return(func(_ context.Context, in model.Order) model.Order {
			return in
		}, nil)

	res, err := s.service.CreateOrder(s.ctx, order)
	s.Require().NoError(err)
	s.Require().Equal(order.UserUUID, res.UserUUID)
	s.Require().Equal(order.PartUUIDs, res.PartUUIDs)
	s.Require().Equal(expectedTotalPrice, res.TotalPrice)
	s.Require().Equal(*order.TransactionUUID, *res.TransactionUUID)
	s.Require().Equal(*order.PaymentMethod, *res.PaymentMethod)
	s.Require().Equal(model.OrderStatusPENDINGPAYMENT, res.Status)
	_, parseErr := uuid.Parse(res.OrderUUID)
	s.Require().NoError(parseErr)
}

func (s *ServiceSuit) TestCreateOrderError(){
	var(
		userUUID = gofakeit.UUID()
		partUUIDs = fakePartsUUIDs(gofakeit.Number(1, 5))
		transactionUUID = gofakeit.UUID()
		paymentMethod = model.PaymentMethod(gofakeit.Number(0, 4))
		orderStatuses = []string{
			string(model.OrderStatusPENDINGPAYMENT),
			string(model.OrderStatusPAID),
			string(model.OrderStatusCANCELLED),
		}
		status = model.OrderStatus(gofakeit.RandomString(orderStatuses))

		order = model.Order{
			UserUUID: userUUID,
			PartUUIDs: partUUIDs,
			TransactionUUID: &transactionUUID,
			PaymentMethod: &paymentMethod,
			Status: status,
		}
		parts = []model.Part{
			{Uuid: partUUIDs[0], Price: gofakeit.Float64Range(1, 100)},
		}
	)
	s.inventoryClient.
		On("GetInventoryModels", s.ctx, partUUIDs).
		Return(parts, nil).
		Once()

	res, err := s.service.CreateOrder(s.ctx, order)
	s.Error(err)
	s.Empty(res)
}

func (s *ServiceSuit) TestCreateOrderRepositoryError(){
	var(
		repoError = gofakeit.Error()
		userUUID = gofakeit.UUID()
		partUUIDs = fakePartsUUIDs(gofakeit.Number(1, 5))
		transactionUUID = gofakeit.UUID()
		paymentMethod = model.PaymentMethod(gofakeit.Number(0, 4))
		orderStatuses = []string{
			string(model.OrderStatusPENDINGPAYMENT),
			string(model.OrderStatusPAID),
			string(model.OrderStatusCANCELLED),
		}
		status = model.OrderStatus(gofakeit.RandomString(orderStatuses))

		order = model.Order{
			UserUUID: userUUID,
			PartUUIDs: partUUIDs,
			TransactionUUID: &transactionUUID,
			PaymentMethod: &paymentMethod,
			Status: status,
		}
		parts = make([]model.Part, 0, len(partUUIDs))
	)
	for _, partUUID := range partUUIDs {
		parts = append(parts, model.Part{
			Uuid: partUUID,
			Price: gofakeit.Float64Range(1, 100),
		})
	}

	s.inventoryClient.
		On("GetInventoryModels", s.ctx, partUUIDs).
		Return(parts, nil).
		Once()

	s.orderRepository.
		On("CreateOrder", s.ctx, mock.MatchedBy(func(in model.Order) bool {
			_, err := uuid.Parse(in.OrderUUID)
			return err == nil && in.Status == model.OrderStatusPENDINGPAYMENT
		})).
		Return(model.Order{}, repoError)

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
