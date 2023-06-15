package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/middlewares"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func BookingRoute(api fiber.Router, bookingRepo *db.BookingRepository) {
	bookingsHandler := handlers.BookingHandler{BookingRepo: bookingRepo}
	bookings := api.Use("/bookings", middlewares.Authorization)
	bookings.Get("/", bookingsHandler.GetBookings)
	bookings.Get("/:id", bookingsHandler.GetBooking)
	bookings.Delete("/:id", bookingsHandler.DeleteBooking)
}
