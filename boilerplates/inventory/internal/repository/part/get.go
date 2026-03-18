package part

import (
	"context"
	"errors"
	"inventory/internal/model"
	repoConverter "inventory/internal/repository/converter"
	repoModel "inventory/internal/repository/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *repository) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	var part repoModel.Part
	err := s.collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&part)
	if err!=nil{
		if errors.Is(err, mongo.ErrNoDocuments){
			return model.Part{}, model.ErrPartNotFound
		}
		return model.Part{}, err
	}

	return repoConverter.PartToModel(&part), nil
}