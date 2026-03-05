package repository

import (
	"context"
	"inventory/internal/model"
)

type InventoryRepository interface {
	// GetPart возвращает деталь по UUID.
	GetPart(ctx context.Context, uuid string) (model.Part, error)
	// ListParts возвращает список деталей с фильтрацией.
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}