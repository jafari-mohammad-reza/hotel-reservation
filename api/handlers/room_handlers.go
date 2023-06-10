package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"time"
)

type RoomHandler struct {
	RoomRepo *db.RoomRepository
}

func (handler *RoomHandler) GetRooms(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rooms, err := handler.RoomRepo.GetAll(ctx)
	if err != nil {
		return err
	}
	if len(rooms) <= 0 {
		data := map[string]interface{}{
			"rooms": []interface{}{},
		}
		return c.JSON(data)
	}
	jsonErr := c.JSON(rooms)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *RoomHandler) GetRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	room, err := handler.RoomRepo.GetById(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}

	jsonErr := c.JSON(room)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *RoomHandler) DeleteRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Params("id")
	err := handler.RoomRepo.Delete(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}
	return nil
}
