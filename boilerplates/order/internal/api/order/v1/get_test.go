package v1

import (
	"order/internal/converter"
	"order/internal/model"
	orderv1 "shared/pkg/openapi/order/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func (s *APISuite) TestGetSuccess(){
    var(
        orderUUID = uuid.New()

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
			OrderUUID: orderUUID.String(),
			UserUUID: userUUID,
			PartUUIDs: converter.UuidsToStrings(partUUIDs),
			TotalPrice: totalPrice,
			TransactionUUID: &transactionUUID,
			PaymentMethod: &paymentMethod,
			Status: status,
		}

        params = orderv1.GetOrderByUUIDParams{
            OrderUUID: orderUUID,
        }

        expectedOrder, _ = converter.OrderToGetOrderByUUIDRes(order)
    )
    s.orderService.On("GetOrder", s.ctx, converter.GetOrderByUUIDParamsToString(params)).Return(order, nil)

    res, err := s.api.GetOrderByUUID(s.ctx, params)

    s.Require().NoError(err)
    s.Require().NotNil(res)
    s.Require().Equal(expectedOrder, res)

}


func (s *APISuite) TestGetNotFoundError(){
    var(
        serviceError = model.ErrOrderNotFound

        orderUUID = uuid.New()
        params = orderv1.GetOrderByUUIDParams{
            OrderUUID: orderUUID,
        }
    )
    s.orderService.On("GetOrder", s.ctx, converter.GetOrderByUUIDParamsToString(params)).Return(model.Order{}, serviceError)

    res, err := s.api.GetOrderByUUID(s.ctx, params)

    s.Require().NoError(err)
    s.Require().IsType(&orderv1.GetOrderByUUIDNotFound{}, res)

    notfound, ok := res.(*orderv1.GetOrderByUUIDNotFound)
    s.Require().True(ok)
    s.Require().Equal(404, notfound.Code)
    s.Require().Equal("order not found", notfound.Message)

}


func (s *APISuite) TestGetInternalError(){
    var(
        serviceError = gofakeit.Error()

        orderUUID = uuid.New()
        params = orderv1.GetOrderByUUIDParams{
            OrderUUID: orderUUID,
        }
    )
    s.orderService.On("GetOrder", s.ctx, converter.GetOrderByUUIDParamsToString(params)).Return(model.Order{}, serviceError)

    res, err := s.api.GetOrderByUUID(s.ctx, params)

    s.Require().NoError(err)
    s.Require().IsType(&orderv1.GetOrderByUUIDInternalServerError{}, res)

    notfound, ok := res.(*orderv1.GetOrderByUUIDInternalServerError)
    s.Require().True(ok)
    s.Require().Equal(500, notfound.Code)
    s.Require().Equal("internal error", notfound.Message)
}