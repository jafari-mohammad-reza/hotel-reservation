package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func ConnectToDB() (*mongo.Client, *mongo.Database) {
	url := os.Getenv("MONGO_URL")
	fmt.Println(url)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		defer func(client *mongo.Client, ctx context.Context) {
			err := client.Disconnect(ctx)
			if err != nil {
				panic(err)
			}
		}(client, context.Background())
	}
	database := client.Database("hotel_reservation")
	return client, database
}
func stringToObjectId(id string) (primitive.ObjectID, error) {
	obi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return obi, nil
}
