package part

import (
	"context"
	"time"
	repoModel "inventory/internal/repository/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct{
	collection *mongo.Collection
}

func NewInventoryRepository(mongoDB *mongo.Database) *repository {
	collection := mongoDB.Collection("parts")

	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "uuid", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateMany(ctx, indexModels)
	if err!= nil{
		panic(err)
	}

	cnt, err := collection.CountDocuments(ctx, bson.M{})
	if err!=nil{
		panic(err)
	}
	
	var parts []repoModel.Part
	if cnt==0{
		parts = seedParts()
		docs := make([]interface{}, 0, len(parts))

		for _, part:=range parts{
			docs = append(docs, part)
		}
		_, err := collection.InsertMany(ctx, docs)
		if err!=nil{
			panic(err)
		}
	}

	return &repository{collection: collection}
}



