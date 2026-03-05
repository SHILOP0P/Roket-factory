package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"order/internal/service/mocks"
)

type APISuite struct{
    ctx context.Context
    orderService *mocks.OrderService
    api *api
    suite.Suite
}

func (s *APISuite) SetupTest(){
    s.ctx = context.Background()
    s.orderService=mocks.NewOrderService(s.T())
    s.api=NewAPI(s.orderService)
}

func TestAPIIntegration(t *testing.T){
    suite.Run(t, new(APISuite))
}
