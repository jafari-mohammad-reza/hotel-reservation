package db

import (
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingRepository struct {
	MongoDbAbstractRepository[types.Booking]
}

func NewBookingRepository(collection *mongo.Collection) *BookingRepository {
	return &BookingRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.Booking]{Collection: collection},
	}
}
