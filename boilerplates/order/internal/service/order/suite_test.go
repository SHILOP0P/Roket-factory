package order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"order/internal/repository/mocks"
)

type ServiceSuit struct{
	suite.Suite
	ctx context.Context
	orderRepository *mocks.OrderRepository
	service *service
}

func(s *ServiceSuit) SetupTest(){
	s.ctx = context.Background()
	s.orderRepository = mocks.NewOrderRepository(s.T())
	s.service = NewOrderService(
		s.orderRepository,
	)
}

func(s *ServiceSuit) TearDownTest(){}

func TestServerIntegration(t *testing.T){
	suite.Run(t, new(ServiceSuit))
}