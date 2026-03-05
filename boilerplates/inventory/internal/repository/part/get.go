package part

import (
	"context"
	"inventory/internal/model"
	repoConverter "inventory/internal/repository/converter"
)

func (s *repository) GetPart(_ context.Context, uuid string) (model.Part, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	part, ok := s.parts[uuid]
	if !ok{
		return model.Part{}, model.ErrPartNotFound
	}

	return repoConverter.PartToModel(part), nil
}