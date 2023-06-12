package db

import (
	"context"
	"errors"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelRepository struct {
	MongoDbAbstractRepository[types.Hotel]
}

func (repo *HotelRepository) CreateHotelByDto(dto *types.CreateHotelDto, ctx context.Context) (*types.Hotel, error) {

	var existHotel types.Hotel
	err := repo.Collection.FindOne(ctx, bson.M{"name": dto.Name}).Decode(&existHotel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Hotel doesn't exist, proceed with creation
			return &types.Hotel{Name: dto.Name, Location: dto.Location}, nil
		}
		// Error other than ErrNoDocuments occurred
		return nil, err
	}
	// Hotel already exists
	return nil, errors.New("hotel already exists with this name")
}

func NewHotelRepository(collection *mongo.Collection) *HotelRepository {
	return &HotelRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.Hotel]{Collection: collection},
	}
}
