package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AbstractRepository interface{}
type MongoDbAbstractRepository struct {
	collection *mongo.Collection
}

func (mongo *MongoDbAbstractRepository) Create(entity any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := mongo.collection.InsertOne(ctx, entity)
	return err
}
func (mongo *MongoDbAbstractRepository) Update(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	obi, err := stringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, updateErr := mongo.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	return updateErr
}
func (mongo *MongoDbAbstractRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	obi, err := stringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, deleteErr := mongo.collection.DeleteOne(ctx, filter)
	return deleteErr
}

func (mongo *MongoDbAbstractRepository) GetAll() (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	entities, getErr := mongo.collection.Find(ctx, nil)
	if getErr != nil {
		return nil, getErr
	}
	return entities, nil
}
func (mongo *MongoDbAbstractRepository) GetById(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	obi, err := stringToObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": obi}
	user := mongo.collection.FindOne(ctx, filter, nil)
	return user, nil
}
