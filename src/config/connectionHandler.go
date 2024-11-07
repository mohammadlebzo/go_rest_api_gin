package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Database
var CTX = context.TODO()

func MakeConnectionMongoDB() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_URL")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(CTX, opts)

	if err != nil {

		panic(err)

	}

	err = client.Ping(CTX, nil)
	MongoClient = client.Database(os.Getenv("DB_NAME"))

	return err
}
