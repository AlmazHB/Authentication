package database

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database

func Init(uri, dbName string) error {
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Fatalf("Error connecting to MongoDB: %s", err.Error())
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		logrus.Fatalf("Error pinging MongoDB: %s", err.Error())
	}

	db = client.Database(dbName)
	return nil
}

func GetDatabase() *mongo.Database {
	return db
}

func Close() error {
	return client.Disconnect(context.Background())
}
