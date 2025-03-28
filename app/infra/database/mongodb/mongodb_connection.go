package mongodb

import (
	"context"

	"lanchonete/bootstrap"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := bootstrap.NewEnv().DBHost
	mongodb_database := bootstrap.NewEnv().DBName

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		println(mongodb_uri)
		println("Error connecting to MongoDB")
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		println("Error connecting to MongoDB 2")
		return nil, err
	}

	return client.Database(mongodb_database), nil
}