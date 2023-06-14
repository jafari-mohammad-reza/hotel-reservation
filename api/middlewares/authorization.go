package middlewares

import (
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/jafari-mohammad-reza/hotel-reservation.git/utils"
)

func Authorization(c *fiber.Ctx) error {
	authorization := c.Get("authorization")
	if len(authorization) <= 0 {
		return &handlers.UnauthorizedError{Message: "unauthorized"}
	}
	splitToken := strings.Split(authorization, "Bearer")
	if len(splitToken) < 2 {
		return &handlers.UnauthorizedError{Message: "unauthorized"}
	}
	token := strings.TrimSpace(splitToken[1])
	userId, err := utils.ExtractPayloadFromJWT(token)
	if err != nil {
		return &handlers.UnauthorizedError{Message: "unauthorized"}
	}
	c.Request().Header.Set("userId", userId)

	return nil
}
