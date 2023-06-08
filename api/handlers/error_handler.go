package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Error interface {
	error
	StatusCode() int
}

type NotFoundError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

type ServerError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}
func (e *BadRequestError) Error() string {
	return e.Message
}

func (e *ServerError) Error() string {
	return e.Message
}

func (e *NotFoundError) StatusCode() int {
	return fiber.StatusNotFound
}
func (e *BadRequestError) StatusCode() int {
	return fiber.StatusBadRequest
}

func (e *ServerError) StatusCode() int {
	return fiber.StatusInternalServerError
}

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(Error); ok {
		return c.Status(e.StatusCode()).JSON(fiber.Map{
			"success": false,
			"message": e.Error(),
		})
	}

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"success": false,
		"message": "Something went wrong",
	})
}
