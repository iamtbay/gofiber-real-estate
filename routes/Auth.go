package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/real-estate-api/controllers"
	"github.com/iamtbay/real-estate-api/middlewares"
)

func Auth() *fiber.App {
	auth := fiber.New()

	auth.Post("/login", middlewares.LogoutMiddleware, controllers.Login)
	auth.Post("/register", middlewares.LogoutMiddleware, controllers.Register)
	//mw
	auth.Post("/logout", middlewares.LoginMiddleware, controllers.Logout)

	return auth

}
