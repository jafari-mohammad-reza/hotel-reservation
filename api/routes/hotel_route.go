package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/middlewares"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func HotelRoute(api fiber.Router, hotelRepo *db.HotelRepository) {
	hotelHandler := handlers.HotelHandler{HotelRepo: hotelRepo}
	hotels := api.Use("/hotels", middlewares.Authorization)
	hotels.Get("/", hotelHandler.GetHotels)
	hotels.Get("/:id", hotelHandler.GetHotel)
	hotels.Post("/", hotelHandler.CreateHotel)
	hotels.Put("/:id", hotelHandler.UpdateHotel)
	hotels.Delete("/:id", hotelHandler.DeleteHotel)
}
