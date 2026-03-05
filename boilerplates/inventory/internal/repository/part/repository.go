package part

import (
	repoModel "inventory/internal/repository/model"
	"sync"
)

type repository struct{
	mu sync.RWMutex
	parts map[string] *repoModel.Part
}

func NewInventoryRepository() *repository {
	return &repository{
		parts: seedParts(),
	}
}



