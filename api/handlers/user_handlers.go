package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"time"
)

type UserHandler struct {
	UserRepo *db.UserRepository
}

func (handler *UserHandler) GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	users, err := handler.UserRepo.GetAll(ctx)
	if err != nil {
		return err
	}
	if len(users) <= 0 {
		data := map[string]interface{}{
			"users": []interface{}{},
		}
		return c.JSON(data)
	}
	jsonErr := c.JSON(users)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user, err := handler.UserRepo.GetById(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}

	jsonErr := c.JSON(user)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}
