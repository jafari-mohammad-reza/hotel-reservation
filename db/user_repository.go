package db

import (
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	MongoDbAbstractRepository[types.User]
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.User]{Collection: collection},
	}
}
