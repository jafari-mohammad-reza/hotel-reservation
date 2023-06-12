package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return &ServerError{Message: jsonErr.Error()}
	}
	return nil
}

func (handler *RoomHandler) CreateRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var dto types.CreateRoomDto
	if err := c.BodyParser(&dto); err != nil {
		return &ServerError{Message: err.Error()}
	}
	room := handler.RoomRepo.CreateRoomFromDto(&dto)
	err := handler.RoomRepo.Create(ctx, room)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}
	return nil
}

func (handler *RoomHandler) UpdateRoom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Params("id")
	var dto types.UpdateRoomDto
	if err := c.BodyParser(&dto); err != nil {
		return &ServerError{Message: err.Error()}
	}
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	dto.ID = hex
	_, err = handler.RoomRepo.UpdateRoomFromDto(&dto, ctx)
	if err != nil {
		return &ServerError{Message: err.Error()}
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
