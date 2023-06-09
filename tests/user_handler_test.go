package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"net/http/httptest"
	"testing"
)

type testDb struct {
	*db.UserRepository
}

func (tdb *testDb) tearDown() error {
	ctx := context.Background()
	if err := tdb.Drop(ctx); err != nil {
		return err
	}
	return nil
}

func setup(t *testing.T) *testDb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	database := client.Database("hotel_reservation_test")
	userRepo := db.NewUserRepository(database.Collection("users"))
	return &testDb{
		UserRepository: userRepo,
	}
}
func TestCreateUser(t *testing.T) {
	testDb := setup(t)
	defer testDb.tearDown()
	t.Failed()
	app := fiber.New()
	userHandler := handlers.UserHandler{UserRepo: testDb.UserRepository}
	dto := types.CreateUserDto{
		Email:    "test@gmail.com",
		Password: "TestPassword",
	}
	app.Post("/", userHandler.CreateUser)
	marshalDto, _ := json.Marshal(dto)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(marshalDto))
	req.Header.Add("Content-Type", "application/json")
	response, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var responseBody types.User
	decodeErr := json.NewDecoder(response.Body).Decode(&responseBody)
	if decodeErr != nil {
		t.Error(decodeErr)
	}
	if responseBody.Email != dto.Email {
		t.Fatal("Response email is not same as dto")
	}
	compare := bcrypt.CompareHashAndPassword([]byte(responseBody.Password), []byte(dto.Password))
	if compare != nil {
		t.Fatal("Response email is not same as dto")
	}
}
