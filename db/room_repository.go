package db

import (
	"context"
	"errors"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository struct {
	MongoDbAbstractRepository[types.Room]
}

func (repo *RoomRepository) CreateRoomFromDto(dto *types.CreateRoomDto) types.Room {
	return types.Room{Type: dto.Type, BasePrice: dto.BasePrice, HotelID: dto.HotelID, Price: dto.BasePrice * float64(dto.Type)}
}
func (repo *RoomRepository) UpdateRoomFromDto(dto *types.UpdateRoomDto, ctx context.Context) (*mongo.UpdateResult, error) {
	var existRoom types.Room
	err := repo.Collection.FindOne(ctx, bson.M{"_id": dto.ID}).Decode(&existRoom)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("there is no room with this id")
		}
		return nil, err
	}

	updateData := bson.M{}
	if dto.Type != nil {
		updateData["type"] = *dto.Type
	}
	if dto.BasePrice != nil {
		updateData["base_price"] = *dto.BasePrice
		if roomType, ok := updateData["type"].(types.RoomType); ok {
			updateData["price"] = *dto.BasePrice * float64(roomType)
		} else {
			updateData["price"] = *dto.BasePrice * float64(existRoom.Type)
		}
	}

	result, updateErr := repo.Collection.UpdateOne(ctx, bson.M{"_id": dto.ID}, bson.M{"$set": updateData})
	if updateErr != nil {
		return nil, updateErr
	}
	return result, err
}

func NewRoomRepository(collection *mongo.Collection) *RoomRepository {
	return &RoomRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.Room]{Collection: collection},
	}
}
