package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/routes"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.CustomErrorHandler,
	})

	app.Get("/notfound", func(c *fiber.Ctx) error {
		return &handlers.NotFoundError{Message: "custom not found error"}
	})

	apiV1 := app.Group("/api/v1")

	_, database := db.ConnectToDB()

	userCollection := database.Collection("users")
	userRepo := db.NewUserRepository(userCollection)
	routes.UserRoute(apiV1, userRepo)

	roomCollection := database.Collection("rooms")
	roomRepo := db.NewRoomRepository(roomCollection)
	routes.RoomRoute(apiV1, roomRepo)

	hotelCollection := database.Collection("hotels")
	hotelRepo := db.NewHotelRepository(hotelCollection)
	routes.HotelRoute(apiV1, hotelRepo)

	if err := app.Listen(":5000"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
