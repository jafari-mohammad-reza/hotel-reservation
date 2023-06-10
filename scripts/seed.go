package main

import (
	"context"
	"fmt"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"time"
)

func seedHotelAndRooms() {
	fmt.Println("Starting seeding hotels and rooms")
	client, database := db.ConnectToDB()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	hotelRepo := db.NewHotelRepository(database.Collection("hotels"))
	roomRepo := db.NewRoomRepository(database.Collection("rooms"))
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, ctx)
	hotels := []interface{}{
		types.Hotel{Name: "Hotel 1", Location: "France"},
		types.Hotel{Name: "Hotel 2", Location: "Germany"},
		types.Hotel{Name: "Hotel 3", Location: "England"},
		types.Hotel{Name: "Hotel 4", Location: "Switzerland"},
	}
	many, hotelsErr := hotelRepo.Collection.InsertMany(ctx, hotels)
	if hotelsErr != nil {
		return
	}
	roomTypes := []types.RoomType{types.DeluxRoomType, types.SingleRoomType, types.SeaSiteRoomType, types.DoubleRoomType}
	for v, _ := range many.InsertedIDs {
		for _, roomType := range roomTypes {
			room := types.Room{
				Type:      roomType,
				HotelID:   many.InsertedIDs[v].(primitive.ObjectID),
				BasePrice: rand.Float64() * 1000,
			}
			room.Price = room.BasePrice * float64(room.Type)
			err := roomRepo.Create(ctx, room)
			if err != nil {
				log.Println("Error creating room:", err)
			}
		}
	}
	fmt.Println("Done seeding hotels and rooms")

}
func main() {
	seedHotelAndRooms()
}
