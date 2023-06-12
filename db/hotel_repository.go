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

			return &types.Hotel{Name: dto.Name, Location: dto.Location}, nil
		}

		return nil, err
	}

	return nil, errors.New("hotel already exists with this name")
}

func (repo *HotelRepository) UpdateHotelByDto(dto *types.UpdateHotelDto, ctx context.Context) (*mongo.UpdateResult, error) {

	var existHotel types.Hotel
	err := repo.Collection.FindOne(ctx, bson.M{"_id": dto.ID}).Decode(&existHotel)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return nil, errors.New("there is no hotel with this id")
		}
		return nil, err
	}

	updateData := bson.M{}
	if dto.Name != nil {
		updateData["name"] = *dto.Name
	}
	if dto.Location != nil {
		updateData["location"] = *dto.Location
	}

	result, updateErr := repo.Collection.UpdateOne(ctx, bson.M{"_id": dto.ID}, bson.M{"$set": updateData})
	if updateErr != nil {
		return nil, updateErr
	}
	return result, err
}

func NewHotelRepository(collection *mongo.Collection) *HotelRepository {
	return &HotelRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.Hotel]{Collection: collection},
	}
}
