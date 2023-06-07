package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	firstName string             `bson:"first_name"`
	lastName  string             `bson:"last_name"`
	email     string             `bson:"email"`
	password  string             `bson:"password"`
}
type UserRepository struct {
	collection *mongo.Collection
}
