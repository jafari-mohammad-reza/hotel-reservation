package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	fmt.Println(headers)
	err := c.Next()
	if err != nil {
		return err
	}
	return nil
}
