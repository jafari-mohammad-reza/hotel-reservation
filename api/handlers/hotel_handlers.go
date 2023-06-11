package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type HotelHandler struct {
	HotelRepo *db.HotelRepository
}

func (handler *HotelHandler) GetHotels(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	lookupStage := bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         "rooms",
			"localField":   "_id",
			"foreignField": "hotelID",
			"as":           "rooms",
		},
	}}

	var hotels []types.Hotel
	aggregateData, err := handler.HotelRepo.Collection.Aggregate(ctx, mongo.Pipeline{lookupStage})

	if err != nil {
		return err
	}
	aggregateErr := aggregateData.All(ctx, &hotels)
	if aggregateErr != nil {
		return aggregateErr
	}

	if len(hotels) <= 0 {
		data := map[string]interface{}{
			"hotels": []interface{}{},
		}
		return c.JSON(data)
	}
	jsonErr := c.JSON(hotels)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *HotelHandler) GetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &ServerError{Message: "Invalid hotel ID"}
	}

	matchStage := bson.D{{
		Key: "$match",
		Value: bson.M{
			"_id": objectID,
		},
	}}

	lookupStage := bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         "rooms",
			"localField":   "_id",
			"foreignField": "hotelID",
			"as":           "rooms",
		},
	}}

	aggregate, aggErr := handler.HotelRepo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage})
	if aggErr != nil {
		return &ServerError{Message: aggErr.Error()}
	}
	var aggregatedHotel []types.Hotel
	aggregateErr := aggregate.All(ctx, &aggregatedHotel)

	if aggregateErr != nil {
		return &ServerError{Message: aggregateErr.Error()}

	}
	jsonErr := c.JSON(aggregatedHotel[0])

	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *HotelHandler) DeleteHotel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Params("id")
	err := handler.HotelRepo.Delete(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}
	return nil
}
