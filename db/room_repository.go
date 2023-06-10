package db

import (
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository struct {
	MongoDbAbstractRepository[types.Room]
}

func NewRoomRepository(collection *mongo.Collection) *RoomRepository {
	return &RoomRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.Room]{Collection: collection},
	}
}
