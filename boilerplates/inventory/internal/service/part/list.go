package part

import (
	"context"
	"inventory/internal/model"
)

func (s *service) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error){
	parts, err := s.InventoryService.ListParts(ctx, filter)
	if err!=nil{
		return nil, err
	}

	return parts, nil
}
