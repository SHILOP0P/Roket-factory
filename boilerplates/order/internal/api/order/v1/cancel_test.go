package v1

import (
	"order/internal/converter"
	"order/internal/model"
	orderv1 "shared/pkg/openapi/order/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

    s.Require().Error(err)
    s.Require().Nil(res)

    st, ok := status.FromError(err)
    s.Require().True(ok)
    s.Require().Equal(codes.NotFound, st.Code())
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

    s.Require().Error(err)
    s.Require().Nil(res)

    st, ok := status.FromError(err)
    s.Require().True(ok)
    s.Require().Equal(codes.Internal, st.Code())
    s.Require().Equal("internal error", st.Message())
}