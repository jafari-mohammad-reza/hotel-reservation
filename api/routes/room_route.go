package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func RoomRoute(api fiber.Router, roomRepo *db.RoomRepository) {
	roomHandler := handlers.RoomHandler{RoomRepo: roomRepo}
	api.Get("/rooms", roomHandler.GetRooms)
	api.Get("/room/:id", roomHandler.GetRoom)
	//api.Post("/rooms/", roomHandler.CreateRoom)
	//api.Put("/room/:id", roomHandler.UodateUser)
	api.Delete("/room/:id", roomHandler.DeleteRoom)
}
