package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/db"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/utils"
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

func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var dto types.CreateUserDto
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	if !types.IsEmailValid(dto.Email) {
		return &BadRequestError{Message: "Invalid email address"}
	}
	if len(dto.Password) < 8 || len(dto.Password) > 16 {
		return &BadRequestError{Message: "Invalid password"}
	}
	createdUser, createUserErr := types.CreateUserFromDto(dto)
	if createUserErr != nil {
		return createUserErr
	}
	err := handler.UserRepo.Create(ctx, createdUser)
	if err != nil {
		return err
	}
	responseErr := c.JSON(createdUser)
	if responseErr != nil {
		return responseErr
	}
	return nil
}

func (handler *UserHandler) LoginUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	var dto types.LoginUserDto
	if err := c.BodyParser(&dto); err != nil {
		return &BadRequestError{Message: err.Error()}
	}
	if !types.IsEmailValid(dto.Email) {
		return &BadRequestError{Message: "Invalid email address"}
	}
	if len(dto.Password) < 8 || len(dto.Password) > 16 {
		return &BadRequestError{Message: "Invalid password"}
	}
	userId, loginErr := handler.UserRepo.Login(ctx, dto.Email, dto.Password)
	if loginErr != nil {
		return &BadRequestError{Message: "invalid credentials"}
	}
	token, tokenErr := utils.GenerateJWTAccessToken(userId)
	if tokenErr != nil {
		return &BadRequestError{Message: tokenErr.Error()}
	}
	jsonBody := map[string]string{
		"token": token,
	}
	jsonErr := c.JSON(jsonBody)
	if jsonErr != nil {
		return &BadRequestError{Message: jsonErr.Error()}
	}
	return nil
}

func (handler *UserHandler) DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := c.Params("id")
	err := handler.UserRepo.Delete(ctx, id)
	if err != nil {
		return &ServerError{Message: err.Error()}
	}
	return nil
}
