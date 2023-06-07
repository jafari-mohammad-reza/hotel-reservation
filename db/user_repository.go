package db

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	MongoDbAbstractRepository
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository{Collection: collection},
	}
}
