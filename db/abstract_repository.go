package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbstractRepository interface{}
type MongoDbAbstractRepository struct {
	Collection *mongo.Collection
}

func (m *MongoDbAbstractRepository) Create(ctx context.Context, entity any) error {

	_, err := m.Collection.InsertOne(ctx, entity)
	return err
}
func (m *MongoDbAbstractRepository) Update(ctx context.Context, id string, update bson.M) error {

	obi, err := stringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, updateErr := m.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	return updateErr
}
func (m *MongoDbAbstractRepository) Delete(ctx context.Context, id string) error {

	obi, err := stringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, deleteErr := m.Collection.DeleteOne(ctx, filter)
	return deleteErr
}

func (m *MongoDbAbstractRepository) GetAll(ctx context.Context) ([]any, error) {

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

func (m *MongoDbAbstractRepository) GetById(ctx context.Context, id string) (any, error) {

	obi, err := stringToObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": obi}
	var data any
	decodeErr := m.Collection.FindOne(ctx, filter, nil).Decode(&data)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return data, nil
}
