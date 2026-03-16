package order

import(
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/mock"

	"order/internal/model"
)

func(s *ServiceSuit) TestPayOrderSuccess(){
	var(
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		transactionUUID = gofakeit.UUID()
		status = model.OrderStatusPAID
		paymentMethod = model.PaymentMethod_PAYMENT_METHOD_CARD

		payInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
			PaymentMethod: &paymentMethod,
			Status: &status,
		}
	)

	s.orderRepository.
		On("GetOrderByUUID", s.ctx, orderUUID).
		Return(model.Order{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
		}, nil).
		Once()

	s.paymentClient.
		On("PayOrder", s.ctx, orderUUID, userUUID, paymentMethod).
		Return(transactionUUID, nil).
		Once()

	s.orderRepository.
		On("UpdateOrder", s.ctx, mock.MatchedBy(func(in model.UpdateOrder) bool {
			return in.OrderUUID == payInfo.OrderUUID &&
				in.UserUUID == payInfo.UserUUID &&
				in.PaymentMethod != nil &&
				*in.PaymentMethod == paymentMethod &&
				in.TransactionUUID != nil &&
				*in.TransactionUUID == transactionUUID
		})).
		Return(nil).
		Once()

	err:= s.service.PayOrder(s.ctx, payInfo)
	s.NoError(err)
}

func (s *ServiceSuit) TestPayOrderErrorPaymentMethodNil(){
	var(
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		payInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
		}
	)

	err := s.service.PayOrder(s.ctx, payInfo)
	s.Error(err)
	s.ErrorIs(err, model.ErrFailPayed)
}

func (s *ServiceSuit) TestPayOrderPaymentClientError(){
	var(
		paymentErr = gofakeit.Error()
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		status = model.OrderStatusPAID
		paymentMethod = model.PaymentMethod_PAYMENT_METHOD_CARD
		payInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
			PaymentMethod: &paymentMethod,
			Status: &status,
		}
	)

	s.orderRepository.
		On("GetOrderByUUID", s.ctx, orderUUID).
		Return(model.Order{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
		}, nil).
		Once()

	s.paymentClient.
		On("PayOrder", s.ctx, orderUUID, userUUID, paymentMethod).
		Return("", paymentErr).
		Once()

	err := s.service.PayOrder(s.ctx, payInfo)
	s.Error(err)
	s.ErrorIs(err, paymentErr)
}

func (s *ServiceSuit) TestPayOrderRepositoryError(){
	var(
		repoErr = gofakeit.Error()
		orderUUID = gofakeit.UUID()
		userUUID = gofakeit.UUID()
		transactionUUID = gofakeit.UUID()
		status = model.OrderStatusPAID
		paymentMethod = model.PaymentMethod_PAYMENT_METHOD_CARD
		payInfo = model.UpdateOrder{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
			PaymentMethod: &paymentMethod,
			Status: &status,
		}
	)

	s.orderRepository.
		On("GetOrderByUUID", s.ctx, orderUUID).
		Return(model.Order{
			OrderUUID: orderUUID,
			UserUUID: userUUID,
		}, nil).
		Once()

	s.paymentClient.
		On("PayOrder", s.ctx, orderUUID, userUUID, paymentMethod).
		Return(transactionUUID, nil).
		Once()

	s.orderRepository.
		On("UpdateOrder", s.ctx, mock.MatchedBy(func(in model.UpdateOrder) bool {
			return in.OrderUUID == payInfo.OrderUUID &&
				in.UserUUID == payInfo.UserUUID &&
				in.TransactionUUID != nil &&
				*in.TransactionUUID == transactionUUID
		})).
		Return(repoErr).
		Once()

	err := s.service.PayOrder(s.ctx, payInfo)
	s.Error(err)
	s.ErrorIs(err, repoErr)
}
