package provider

import (
	"app/pkg/example"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func ProvideExample(ctx context.Context, db *gorm.DB, mg *mongo.Client) (*example.Example, error) {
	return example.New(ctx, example.Config{
		GormDB:      db,
		MongoClient: mg,
	})
}
