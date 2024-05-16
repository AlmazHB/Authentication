package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	Users  *mongo.Collection
)

func Init(uri, database, collection string) error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	localclient, err := mongo.Connect(context.Background(), opts)
	if err != nil {

		return err
	}
	client = localclient

	Users = client.Database(database).Collection(collection)
	err = client.Database(database).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	if err != nil {

		return err
	}
	fmt.Println("MongoDB connected soccessfuly!")

	return nil
}
func Close() error {
	return client.Disconnect(context.Background())
}
