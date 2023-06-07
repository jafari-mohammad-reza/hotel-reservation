package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/routes"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

func main() {
	app := fiber.New()
	go func() {
		db.ConnectToDB()
		fmt.Println("Connected to mongodb")
	}()
	apiV1 := app.Group("/api/v1")
	routes.UserRoute(apiV1)
	listenErr := app.Listen(":5000")
	if listenErr != nil {
		return
	}
}
