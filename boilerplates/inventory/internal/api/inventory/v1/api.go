package v1

import (
	inventoryV1 "shared/pkg/proto/inventory/v1"
	"inventory/internal/service"
)

type api struct {
	inventoryV1.UnimplementedInventoryServiceServer

	inventoryService	service.InventoryService
}

func NewAPI(inventoryService service.InventoryService) *api{
	return &api{
		inventoryService: inventoryService,
	}
}