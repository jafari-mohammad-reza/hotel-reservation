package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
)

func UserRoute(api fiber.Router) {
	api.Get("/users", handlers.GetUsers)
}
