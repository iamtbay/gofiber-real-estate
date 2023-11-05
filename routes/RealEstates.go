package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/real-estate-api/controllers"
	"github.com/iamtbay/real-estate-api/middlewares"
)

func RealEstates() *fiber.App {
	r := fiber.New()

	r.Post("/q", controllers.GetAllEstates)
	r.Post("/search", controllers.GetEstatesByQuery)

	//r.Use("/upload", middlewares.LoginMiddleware, controllers.UploadFile)
	r.Post("/upload", middlewares.LoginMiddleware, controllers.UploadFile)
	r.Post("/", controllers.AddNewEstate)

	//single estate opertationss
	singleEstate := r.Group("/estate")
	singleEstate.Get("/:id", controllers.GetSingleEstate)
	//mw for protection
	singleEstate.Use(middlewares.LoginMiddleware)
	singleEstate.Patch("/:id", controllers.UpdateEstate)
	singleEstate.Delete("/:id", controllers.DeleteEstate)
	//

	return r

}
