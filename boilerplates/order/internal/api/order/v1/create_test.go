package v1

import (
	"order/internal/converter"
	"order/internal/model"
	orderv1 "shared/pkg/openapi/order/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func (s *APISuite) TestCreateSuccess(){
    var(
        userUUID = uuid.New()
        listPartsUUIDs = fakePartsUUIDs(gofakeit.Number(1, 5))

        params = orderv1.CreateOrderRequest{
            UserUUID: userUUID,
            PartUuids: listPartsUUIDs,
        }

        orderUUID = gofakeit.UUID()
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
			UserUUID: userUUID.String(),
			PartUUIDs: converter.UuidsToStrings(partUUIDs),
			TotalPrice: totalPrice,
			TransactionUUID: &transactionUUID,
			PaymentMethod: &paymentMethod,
			Status: status,
		}

        exeptedOrder, _ = converter.OrderToCreateOrderResponse(order)
    )
    s.orderService.On("CreateOrder", s.ctx, converter.CreateOrderRequestToModel(&params)).Return(order, nil)

    res, err := s.api.CreateOrder(s.ctx, &params)

    s.Require().NoError(err)
    s.Require().NotNil(res)
    s.Require().Equal(exeptedOrder, res)

}


func (s *APISuite) TestCreateFailCreatedError(){
    var(
        userUUID = uuid.New()
        listPartsUUIDs = fakePartsUUIDs(gofakeit.Number(1, 5))

        serviceError = model.ErrFailCreated

        params = orderv1.CreateOrderRequest{
            UserUUID: userUUID,
            PartUuids: listPartsUUIDs,
        }
    )
    s.orderService.On("CreateOrder", s.ctx, converter.CreateOrderRequestToModel(&params)).Return(model.Order{}, serviceError)

    res, err := s.api.CreateOrder(s.ctx, &params)
    s.Require().NoError(err)
    s.Require().NotNil(res)

    internalErr, ok := res.(*orderv1.CreateOrderInternalServerError)
    s.Require().True(ok)
    s.Require().Equal(500, internalErr.Code)
    s.Require().Equal(serviceError.Error(), internalErr.Message)
}



func fakePartsUUIDs(n int)[]uuid.UUID{
	listPartsUUIDs := make([]uuid.UUID, n)

	for i := 0; i < n; i++ {
		partUUID := uuid.New()
		listPartsUUIDs[i] = partUUID
	}
	return listPartsUUIDs
}
