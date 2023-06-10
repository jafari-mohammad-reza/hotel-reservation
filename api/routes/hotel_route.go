package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func HotelRoute(api fiber.Router, hotelRepo *db.HotelRepository) {
	hotelHandler := handlers.HotelHandler{HotelRepo: hotelRepo}
	api.Get("/hotels", hotelHandler.GetHotels)
	api.Get("/hotels/:id", hotelHandler.GetHotel)
	api.Delete("/hotels/:id", hotelHandler.DeleteHotel)
}
