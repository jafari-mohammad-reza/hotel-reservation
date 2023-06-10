package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"time"
)

type HotelHandler struct {
	HotelRepo *db.HotelRepository
}

func (handler *HotelHandler) GetHotels(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	hotels, err := handler.HotelRepo.GetAll(ctx)
	if err != nil {
		return err
	}
	if len(hotels) <= 0 {
		data := map[string]interface{}{
			"hotels": []interface{}{},
		}
		return c.JSON(data)
	}
	jsonErr := c.JSON(hotels)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *HotelHandler) GetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user, err := handler.HotelRepo.GetById(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}

	jsonErr := c.JSON(user)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (handler *HotelHandler) DeleteHotel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Params("id")
	err := handler.HotelRepo.Delete(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}
	return nil
}
