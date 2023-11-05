package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes() *fiber.App {
	app := fiber.New()

	// app.Get("/images/:id", func(c *fiber.Ctx) error {

	// })
	//Auth routes
	app.Mount("/auth", Auth())
	//Real estate routes
	app.Mount("/estates", RealEstates())

	//

	return app
}
