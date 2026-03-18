package part

import (
	"context"
	"errors"
	"inventory/internal/model"
	repoConverter "inventory/internal/repository/converter"
	repoModel "inventory/internal/repository/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *repository) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	mongoFilter:= buildMongoFilter(repoConverter.PartsFilterToRepoModel(filter))
	
	cursor, err := s.collection.Find(ctx, mongoFilter)
	if err!=nil{
		if errors.Is(err, mongo.ErrNoDocuments){
			return nil, model.ErrPartNotFound
		}
		return nil, err
	}
	defer func(){
		cerr:=cursor.Close(ctx)
		if cerr!=nil{
			log.Printf("faled close cursor: %v", cerr)
		}
	}()
	var repoParts []repoModel.Part
	err = cursor.All(ctx, &repoParts)
	if err!=nil{
		return nil, err
	}

	return repoConverter.PartsToModel(repoParts), nil
}


func buildMongoFilter(filter repoModel.PartsFilter) bson.M {
	mongoFilter := bson.M{}
	if len(filter.Uuids)>0{
		mongoFilter["uuid"] = bson.M{"$in": filter.Uuids}
	}
	if len(filter.Names)>0{
		mongoFilter["name"] = bson.M{"$in": filter.Names}
	}
	if len(filter.Categories)>0{
		mongoFilter["category"] = bson.M{"$in": filter.Categories}
	}
	if len(filter.Tags)>0{
		mongoFilter["tags"] = bson.M{"$in": filter.Tags}
	}
	if len(filter.ManufacturerCountries)>0{
		mongoFilter["manufacturer.country"] = bson.M{"$in":filter.ManufacturerCountries}
	}
	return mongoFilter
}