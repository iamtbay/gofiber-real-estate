package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/real-estate-api/routes"
)

func StartServer() {
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Real-Estate-Fiber",
		AppName:       "iamtbay-real-estate",
	})
	//
	app.Mount("/api/v1", routes.Routes())
	app.Static(
		"/images",
		"./public/uploads",
	)
	//

	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
