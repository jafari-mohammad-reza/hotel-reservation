package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/routes"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

func main() {
	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	var wt sync.WaitGroup
	wt.Add(1)
	var _ *mongo.Client
	var database *mongo.Database
	go func() {
		_, database = db.ConnectToDB()
		fmt.Println("Connected to mongodb")
		wt.Done()
	}()
	wt.Wait()
	userCollection := database.Collection("users")
	userRepo := db.NewUserRepository(userCollection)
	routes.UserRoute(apiV1, userRepo)
	listenErr := app.Listen(":5000")
	if listenErr != nil {
		return
	}
}
