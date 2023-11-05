package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/real-estate-api/helpers"
)

// check user. if logged in let him go if not send an error
func LoginMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("accessToken")
	url := c.Request().URI()
	if fmt.Sprint(url) != "http://localhost:8080/api/v1/auth/logout" {
		_, err := helpers.ParseJWT(accessToken)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
		}
	}

	if accessToken == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden!"})
	}
	return c.Next()
}

// check user. if logged in send an error if not let him go.
func LogoutMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("accessToken")
	if accessToken == "" {
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden!"})
}
