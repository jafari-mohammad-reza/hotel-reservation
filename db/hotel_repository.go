package db

import (
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelRepository struct {
	MongoDbAbstractRepository[types.Hotel]
}

func NewHotelRepository(collection *mongo.Collection) *HotelRepository {
	return &HotelRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.Hotel]{Collection: collection},
	}
}
