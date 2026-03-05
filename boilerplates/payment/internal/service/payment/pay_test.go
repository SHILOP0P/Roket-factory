package payment

import (
	"payment/internal/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	var (
		orderUUID  = gofakeit.UUID()
		userUUID   = gofakeit.UUID()
		payMethod  = model.PaymentMethod(gofakeit.Number(0, 4))
	)

	txID, err := s.service.PayOrder(
		s.ctx,
		orderUUID,
		userUUID,
		payMethod,
	)

	s.Require().NoError(err)
	s.Require().NotEmpty(txID)
	_, parseErr := uuid.Parse(txID)
	s.Require().NoError(parseErr)
}

func (s *ServiceSuite) TestPayOrderReturnsDifferentTransactions() {
	var (
		orderUUID = gofakeit.UUID()
		userUUID  = gofakeit.UUID()
		payMethod = model.PaymentMethod(gofakeit.Number(0, 4))
	)

	txID1, err1 := s.service.PayOrder(
		s.ctx,
		orderUUID,
		userUUID,
		payMethod,
	)
	txID2, err2 := s.service.PayOrder(
		s.ctx,
		orderUUID,
		userUUID,
		payMethod,
	)

	s.Require().NoError(err1)
	s.Require().NoError(err2)
	s.Require().NotEqual(txID1, txID2)
}
