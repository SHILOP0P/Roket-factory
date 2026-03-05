package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"payment/internal/service/mocks"
)

type APISuite struct{
    ctx context.Context
    paymentService *mocks.PaymentService
    api *api
    suite.Suite
}

func (s *APISuite) SetupTest(){
    s.ctx = context.Background()
    s.paymentService=mocks.NewPaymentService(s.T())
    s.api=NewAPI(s.paymentService)
}

func TestAPIIntegration(t *testing.T){
    suite.Run(t, new(APISuite))
}
