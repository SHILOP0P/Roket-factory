package v1

import (
    "context"
    "testing"
    "github.com/stretchr/testify/suite"
    "inventory/internal/service/mocks"
)

type APISuite struct{
    suite.Suite
    ctx context.Context
    inventoryService *mocks.InventoryService
    api *api
}

func (s *APISuite) SetupTest(){
    s.ctx = context.Background()
    s.inventoryService = mocks.NewInventoryService(s.T())
    s.api = NewAPI(
        s.inventoryService,
    )
}

func (s *APISuite) TearDownTest(){
}

func TestApiIntegration(t *testing.T){
    suite.Run(t, new(APISuite))
}