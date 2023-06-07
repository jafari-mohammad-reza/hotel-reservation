package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
)

type UserHandler struct {
	UserRepo *db.UserRepository
}

func (handler *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := handler.UserRepo.GetAll()
	if err != nil {
		return err
	}

	jsonErr := c.JSON(users)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}
