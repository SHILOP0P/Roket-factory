package v1

import(
	inventory_v1 "shared/pkg/proto/inventory/v1"
	"order/internal/service"
)

type inventoryClient struct{
	client inventory_v1.InventoryServiceClient
}

func NewInventoryClient(client inventory_v1.InventoryServiceClient) service.InventoryClient{
	return &inventoryClient{
		client: client,
	}
}