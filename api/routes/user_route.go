package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func UserRoute(api fiber.Router, userRepo *db.UserRepository) {
	userHandler := handlers.UserHandler{UserRepo: userRepo}
	api.Get("/users", userHandler.GetUsers)
}
