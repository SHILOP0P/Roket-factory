package part

import (
	"inventory/internal/repository"
)


type service struct {
	InventoryService repository.InventoryRepository
}

func NewService(inventoryService repository.InventoryRepository) *service {
	return &service{
		InventoryService: inventoryService,
	}
}