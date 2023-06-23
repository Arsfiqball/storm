package provider

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProvideMongo(ctx context.Context, cfg *viper.Viper) (*mongo.Client, error) { // Use the SetServerAPIOptions() method to set the Stable API version to 1
	opts := options.
		Client().
		ApplyURI(cfg.GetString("mongo_uri"))

	return mongo.Connect(ctx, opts)
}
