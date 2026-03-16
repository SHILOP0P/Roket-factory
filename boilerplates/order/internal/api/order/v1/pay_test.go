package v1

import (
	"order/internal/converter"
	"order/internal/model"
	orderv1 "shared/pkg/openapi/order/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func (s *APISuite) TestPaySuccess(){
    var(
        orderUUID = uuid.New()

        methods = []orderv1.PaymentMethod{
            orderv1.PaymentMethodPAYMENTMETHODUNSPECIFIED,
            orderv1.PaymentMethodPAYMENTMETHODCARD,
            orderv1.PaymentMethodPAYMENTMETHODSBP,
            orderv1.PaymentMethodPAYMENTMETHODCREDITCARD,
            orderv1.PaymentMethodPAYMENTMETHODINVESTORMONEY,
        }
        paymentMethodPay = methods[gofakeit.Number(0, len(methods)-1)]

        req = &orderv1.PayOrderRequest{
            PaymentMethod: paymentMethodPay,
        }

        paramsPay = orderv1.PayOrderParams{
            OrderUUID: orderUUID,
        }


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

        paramsGet = orderv1.GetOrderByUUIDParams{
            OrderUUID: orderUUID,
        }
    )

    s.orderService.On("GetOrder", s.ctx, converter.GetOrderByUUIDParamsToString(paramsGet)).Return(order, nil).Once()
    update := converter.PayOrderRequestToModel(paramsPay, req)
    update.UserUUID = userUUID
    s.orderService.On("PayOrder", s.ctx, update).Return(nil).Once()
    s.orderService.On("GetOrder", s.ctx, converter.GetOrderByUUIDParamsToString(paramsGet)).Return(order, nil).Once()

    res, err := s.api.PayOrder(s.ctx, req, paramsPay)

    s.Require().NoError(err)
    okRes, ok := res.(*orderv1.PayOrderResponse)
    s.Require().True(ok)
    s.Require().Equal(uuid.MustParse(transactionUUID), okRes.TransactionUUID)


}


func (s *APISuite) TestPayNotFoundError(){
	var(
		serviceError = model.ErrOrderNotFound
		orderUUID = uuid.New()

		req = &orderv1.PayOrderRequest{
			PaymentMethod: orderv1.PaymentMethodPAYMENTMETHODCARD,
		}
		paramsPay = orderv1.PayOrderParams{
			OrderUUID: orderUUID,
		}
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.Order{}, serviceError).Once()

	res, err := s.api.PayOrder(s.ctx, req, paramsPay)

	s.Require().NoError(err)
	s.Require().IsType(&orderv1.PayOrderNotFound{}, res)

	notFound, ok := res.(*orderv1.PayOrderNotFound)
	s.Require().True(ok)
	s.Require().Equal(404, notFound.Code)
	s.Require().Equal("order not found", notFound.Message)
}

func (s *APISuite) TestPayInternalError(){
	var(
		serviceError = gofakeit.Error()
		orderUUID = uuid.New()
		userUUID = gofakeit.UUID()

		req = &orderv1.PayOrderRequest{
			PaymentMethod: orderv1.PaymentMethodPAYMENTMETHODCARD,
		}
		paramsPay = orderv1.PayOrderParams{
			OrderUUID: orderUUID,
		}
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.Order{
		OrderUUID: orderUUID.String(),
		UserUUID: userUUID,
	}, nil).Once()
	update := converter.PayOrderRequestToModel(paramsPay, req)
	update.UserUUID = userUUID
	s.orderService.On("PayOrder", s.ctx, update).Return(serviceError).Once()

	res, err := s.api.PayOrder(s.ctx, req, paramsPay)

	s.Require().NoError(err)
	s.Require().IsType(&orderv1.PayOrderInternalServerError{}, res)

	internalErr, ok := res.(*orderv1.PayOrderInternalServerError)
	s.Require().True(ok)
	s.Require().Equal(500, internalErr.Code)
	s.Require().Equal("internal error", internalErr.Message)
}

func (s *APISuite) TestPayFailPayedError(){
	var(
		serviceError = model.ErrFailPayed
		orderUUID = uuid.New()

		req = &orderv1.PayOrderRequest{
			PaymentMethod: orderv1.PaymentMethodPAYMENTMETHODCARD,
		}
		paramsPay = orderv1.PayOrderParams{
			OrderUUID: orderUUID,
		}
	)

	userUUID := gofakeit.UUID()
	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.Order{
		OrderUUID: orderUUID.String(),
		UserUUID: userUUID,
	}, nil).Once()
	update := converter.PayOrderRequestToModel(paramsPay, req)
	update.UserUUID = userUUID
	s.orderService.On("PayOrder", s.ctx, update).Return(serviceError).Once()

	res, err := s.api.PayOrder(s.ctx, req, paramsPay)

	s.Require().NoError(err)
	s.Require().IsType(&orderv1.PayOrderConflict{}, res)

	conflictErr, ok := res.(*orderv1.PayOrderConflict)
	s.Require().True(ok)
	s.Require().Equal(409, conflictErr.Code)
	s.Require().Equal("order cannot be paid", conflictErr.Message)
}

func (s *APISuite) TestPayGetOrderNotFoundError(){
	var(
		orderUUID = uuid.New()
		serviceError = model.ErrOrderNotFound

		req = &orderv1.PayOrderRequest{
			PaymentMethod: orderv1.PaymentMethodPAYMENTMETHODCARD,
		}
		paramsPay = orderv1.PayOrderParams{
			OrderUUID: orderUUID,
		}
	)

	userUUID := gofakeit.UUID()
	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.Order{
		OrderUUID: orderUUID.String(),
		UserUUID: userUUID,
	}, nil).Once()
	update := converter.PayOrderRequestToModel(paramsPay, req)
	update.UserUUID = userUUID
	s.orderService.On("PayOrder", s.ctx, update).Return(nil).Once()
	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.Order{}, serviceError).Once()

	res, err := s.api.PayOrder(s.ctx, req, paramsPay)

	s.Require().NoError(err)
	s.Require().IsType(&orderv1.PayOrderNotFound{}, res)

	notFound, ok := res.(*orderv1.PayOrderNotFound)
	s.Require().True(ok)
	s.Require().Equal(404, notFound.Code)
	s.Require().Equal("order not found", notFound.Message)
}

func (s *APISuite) TestPayGetOrderInternalError(){
	var(
		orderUUID = uuid.New()
		serviceError = gofakeit.Error()

		req = &orderv1.PayOrderRequest{
			PaymentMethod: orderv1.PaymentMethodPAYMENTMETHODCARD,
		}
		paramsPay = orderv1.PayOrderParams{
			OrderUUID: orderUUID,
		}
	)

	userUUID := gofakeit.UUID()
	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.Order{
		OrderUUID: orderUUID.String(),
		UserUUID: userUUID,
	}, nil).Once()
	update := converter.PayOrderRequestToModel(paramsPay, req)
	update.UserUUID = userUUID
	s.orderService.On("PayOrder", s.ctx, update).Return(nil).Once()
	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.Order{}, serviceError).Once()

	res, err := s.api.PayOrder(s.ctx, req, paramsPay)

	s.Require().NoError(err)
	s.Require().IsType(&orderv1.PayOrderInternalServerError{}, res)

	internalErr, ok := res.(*orderv1.PayOrderInternalServerError)
	s.Require().True(ok)
	s.Require().Equal(500, internalErr.Code)
	s.Require().Equal("internal error", internalErr.Message)
}
