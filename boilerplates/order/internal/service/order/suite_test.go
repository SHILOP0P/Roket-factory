package order

import (
	"context"
	"testing"

	serviceMocks "order/internal/service/mocks"
	"github.com/stretchr/testify/suite"
	"order/internal/repository/mocks"
)

type ServiceSuit struct{
	suite.Suite
	ctx context.Context
	orderRepository *mocks.OrderRepository
	service *service
	inventoryClient *serviceMocks.InventoryClient
	paymentClient *serviceMocks.PaymentClient
}

func(s *ServiceSuit) SetupTest(){
	s.ctx = context.Background()
	s.orderRepository = mocks.NewOrderRepository(s.T())
	s.inventoryClient = serviceMocks.NewInventoryClient(s.T())
	s.paymentClient = serviceMocks.NewPaymentClient(s.T())
	
	s.service = NewOrderService(
		s.orderRepository, s.inventoryClient, s.paymentClient,
	)
}

func(s *ServiceSuit) TearDownTest(){}

func TestServerIntegration(t *testing.T){
	suite.Run(t, new(ServiceSuit))
}
