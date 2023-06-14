package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/middlewares"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func RoomRoute(api fiber.Router, roomRepo *db.RoomRepository) {
	roomHandler := handlers.RoomHandler{RoomRepo: roomRepo}
	rooms := api.Use("/rooms", middlewares.Authorization)
	rooms.Get("/", roomHandler.GetRooms)
	rooms.Post("/", roomHandler.CreateRoom)
	rooms.Get("/:id", roomHandler.GetRoom)
	rooms.Put("/:id", roomHandler.UpdateRoom)
	rooms.Delete("/:id", roomHandler.DeleteRoom)
}
