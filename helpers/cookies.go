package helpers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CREATE COOKIE
func CreateCookie(name, token string) fiber.Handler {
	fmt.Println("hello from handler")
	return func(c *fiber.Ctx) error {
		fmt.Println("hello inside f")
		c.Cookie(&fiber.Cookie{
			Name:     name,
			Value:    token,
			Path:     "/",
			MaxAge:   1,
			Expires:  time.Now().Add(time.Hour * 12),
			HTTPOnly: true,
		})
		return nil
	}

}

// DELETE COOKIE
func DeleteCookie(name string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:   name,
			MaxAge: -1,
		})
		return nil
	}
}
