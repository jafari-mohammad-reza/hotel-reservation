package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func RoomRoute(api fiber.Router, roomRepo *db.RoomRepository) {
	roomHandler := handlers.RoomHandler{RoomRepo: roomRepo}
	api.Get("/rooms", roomHandler.GetRooms)
	api.Get("/rooms/:id", roomHandler.GetRoom)
	api.Post("/rooms/", roomHandler.CreateRoom)
	api.Put("/rooms/:id", roomHandler.UpdateRoom)
	api.Delete("/rooms/:id", roomHandler.DeleteRoom)
}
