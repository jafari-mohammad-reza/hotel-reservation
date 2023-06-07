package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AbstractRepository interface{}
type MongoDbAbstractRepository struct {
	Collection *mongo.Collection
}

func (m *MongoDbAbstractRepository) Create(entity any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.InsertOne(ctx, entity)
	return err
}
func (m *MongoDbAbstractRepository) Update(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	obi, err := stringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, updateErr := m.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	return updateErr
}
func (m *MongoDbAbstractRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	obi, err := stringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, deleteErr := m.Collection.DeleteOne(ctx, filter)
	return deleteErr
}

func (m *MongoDbAbstractRepository) GetAll() ([]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var entities []any
	err = cursor.All(ctx, &entities)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []any{}, nil
		}
		return nil, err
	}
	return entities, nil
}

func (m *MongoDbAbstractRepository) GetById(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	obi, err := stringToObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": obi}
	user := m.Collection.FindOne(ctx, filter, nil)
	return user, nil
}
