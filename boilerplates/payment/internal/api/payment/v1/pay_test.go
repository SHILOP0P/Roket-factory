package v1

import (
	"payment/internal/converter"
	payment "shared/pkg/proto/payment/v1"

	"github.com/brianvoe/gofakeit/v6"
)

func (s *APISuite) TestPaySuccess(){
    var(
        orderUUID = gofakeit.UUID()
        userUUID = gofakeit.UUID()

        transactionUUID = gofakeit.UUID()

        methods = []payment.PaymentMethod{
            payment.PaymentMethod_PAYMENT_METHOD_UNSPECIFIED,
            payment.PaymentMethod_PAYMENT_METHOD_CARD,
            payment.PaymentMethod_PAYMENT_METHOD_SBP,
            payment.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD,
            payment.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY,
        }
        paymentMethodPay = methods[gofakeit.Number(0, len(methods)-1)]

        req = payment.PayOrderRequest{
            OrderUuid: orderUUID,
            UserUuid: userUUID,
            PaymentMethod: paymentMethodPay,
        }
    )

    s.paymentService.On("PayOrder", s.ctx, orderUUID, userUUID, converter.PaymentMethodProtoToModel(paymentMethodPay)).Return(transactionUUID, nil)

    res, err := s.api.PayOrder(s.ctx, &req)

    s.Require().NoError(err)
    s.Require().IsType(&payment.PayOrderResponse{}, res)
    s.Require().Equal(transactionUUID, res.TransactionUuid)

}