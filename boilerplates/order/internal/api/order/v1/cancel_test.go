package v1

import (
	"order/internal/converter"
	"order/internal/model"
	orderv1 "shared/pkg/openapi/order/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func (s *APISuite) TestCancelSuccess(){
    var(
        orderUUID = uuid.New()

        params = orderv1.CancelOrderParams{
            OrderUUID: orderUUID,
        }
    )
    s.orderService.On("CancelOrder", s.ctx, converter.CancelOrderParamsToModel(params)).Return(nil)

    res, err := s.api.CancelOrder(s.ctx, params)

    s.Require().NoError(err)
    s.Require().NotNil(res)
    s.Require().IsType(&orderv1.CancelOrderNoContent{}, res)

}


func (s *APISuite) TestCancelNotFoundError(){
    var(
        serviceError = model.ErrOrderNotFound

        orderUUID = uuid.New()
        params = orderv1.CancelOrderParams{
            OrderUUID: orderUUID,
        }
    )
    s.orderService.On("CancelOrder", s.ctx, converter.CancelOrderParamsToModel(params)).Return(serviceError)

	res, err := s.api.CancelOrder(s.ctx, params)

	s.Require().NoError(err)
	s.Require().IsType(&orderv1.CancelOrderNotFound{}, res)

	notFound, ok := res.(*orderv1.CancelOrderNotFound)
	s.Require().True(ok)
	s.Require().Equal(404, notFound.Code)
	s.Require().Equal("Order not found", notFound.Message)
}

func (s *APISuite) TestCancelInternalError(){
    var(
        serviceError = gofakeit.Error()

        orderUUID = uuid.New()
        params = orderv1.CancelOrderParams{
            OrderUUID: orderUUID,
        }
    )
    s.orderService.On("CancelOrder", s.ctx, converter.CancelOrderParamsToModel(params)).Return(serviceError)

	res, err := s.api.CancelOrder(s.ctx, params)

	s.Require().NoError(err)
	s.Require().IsType(&orderv1.CancelOrderInternalServerError{}, res)

	internalErr, ok := res.(*orderv1.CancelOrderInternalServerError)
	s.Require().True(ok)
	s.Require().Equal(500, internalErr.Code)
	s.Require().Equal("internal error", internalErr.Message)
}
