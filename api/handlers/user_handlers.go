package handlers

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	users := map[string]string{"msg": "user handler"}
	jsonErr := c.JSON(users)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}
