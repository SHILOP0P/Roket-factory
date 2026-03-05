package part

import (
	"context"
	"inventory/internal/model"
)

func (s *service) GetPart(ctx context.Context, uuid string) (model.Part, error){
	part, err := s.InventoryService.GetPart(ctx, uuid)
	if err!=nil{
		return model.Part{}, err
	}

	return part, nil
}
