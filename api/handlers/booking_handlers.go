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

type BookingHandler struct {
	BookingRepo *db.BookingRepository
}

func (handler *BookingHandler) GetBookings(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	lookupRoom := bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         "rooms",
			"localField":   "roomID",
			"foreignField": "_id",
			"as":           "room",
		},
	}}
	lookupUser := bson.D{{
		Key: "$lookup",
		Value: bson.M{
			"from":         "users",
			"localField":   "userID",
			"foreignField": "_id",
			"as":           "user",
		},
	}}

	var bookings []types.Booking
	aggregateData, err := handler.BookingRepo.Collection.Aggregate(ctx, mongo.Pipeline{lookupRoom, lookupUser})

	if err != nil {
		return &ServerError{Message: err.Error()}
	}
	aggregateErr := aggregateData.All(ctx, &bookings)
	if aggregateErr != nil {
		return &ServerError{Message: aggregateErr.Error()}
	}

	if len(bookings) <= 0 {
		data := map[string]interface{}{
			"bookings": []interface{}{},
		}
		return c.JSON(data)
	}
	jsonErr := c.JSON(bookings)
	if jsonErr != nil {
		return &ServerError{Message: jsonErr.Error()}
	}
	return nil
}

func (handler *BookingHandler) GetBooking(c *fiber.Ctx) error {
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

	aggregate, aggErr := handler.BookingRepo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage})
	if aggErr != nil {
		return &ServerError{Message: aggErr.Error()}
	}
	var aggregatedBooking []types.Booking
	aggregateErr := aggregate.All(ctx, &aggregatedBooking)

	if aggregateErr != nil {
		return &ServerError{Message: aggregateErr.Error()}

	}
	jsonErr := c.JSON(aggregatedBooking[0])

	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *BookingHandler) DeleteBooking(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Params("id")
	err := handler.BookingRepo.Delete(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}
	return nil
}
