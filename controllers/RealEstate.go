package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/real-estate-api/database"
	"github.com/iamtbay/real-estate-api/helpers"
	"github.com/iamtbay/real-estate-api/models"
)

var estateDB = database.InitEstates()

// GET ALL ESTATES
func GetEstatesByQuery(c *fiber.Ctx) error {
	// get page
	var pageQuery = c.Query("page")
	pageNum, _ := strconv.Atoi(pageQuery)
	// get other inputs for search
	var searchInputs *models.EstateSearch
	if err := c.BodyParser(&searchInputs); err != nil {
		log.Fatal(err)
	}

	//
	allEstates, err := estateDB.GetEstatesByQuery(pageNum, searchInputs)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":          "successful",
		"totalEstateCount": allEstates.TotalEstateCount,
		"totalPage":        allEstates.TotalPage,
		"currentPage":      allEstates.CurrentPage,
		"zdata":            allEstates.Data,
	})
}

func GetAllEstates(c *fiber.Ctx) error {
	var pageQuery = c.Query("page")
	pageNum, _ := strconv.Atoi(pageQuery)
	allEstates, err := estateDB.GetAllEstates(pageNum)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":          "successful",
		"totalEstateCount": allEstates.TotalEstateCount,
		"totalPage":        allEstates.TotalPage,
		"currentPage":      allEstates.CurrentPage,
		"zdata":            allEstates.Data,
	})
}

// ADD NEW ESTATE
func AddNewEstate(c *fiber.Ctx) error {
	//get token
	accessToken := c.Cookies("accessToken")
	//get user id
	userID, err := helpers.ParseJWT(accessToken)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
	}

	//
	var estateInfo *models.Estate

	//
	if err := c.BodyParser(&estateInfo); err != nil {
		log.Fatal(err)
	}
	estateInfo.OwnerID = userID
	estateInfo.CreatedAt = time.Now()
	//
	err = estateDB.AddNewEstate(estateInfo)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "succesfully inserted!",
	})
}

// GET SINGLE ESTATE
func GetSingleEstate(c *fiber.Ctx) error {
	idParams := c.Params("id")

	estate, err := estateDB.GetSingleEstate(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": estate,
	})
}

// UPDATE ESTATE
func UpdateEstate(c *fiber.Ctx) error {
	//get token
	accessToken := c.Cookies("accessToken")
	//get user id
	userID, err := helpers.ParseJWT(accessToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	idString := c.Params("id")
	var estateInfo *models.Estate
	if err := c.BodyParser(&estateInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//first update estate
	deleteImgs, err := estateDB.UpdateEstate(estateInfo, idString, userID)
	// delete inequal images on sv.
	if len(deleteImgs) > 0 {
		err := helpers.ImageDeleter(deleteImgs)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succesfully updated",
	})

}

// DELETE ESTATE
func DeleteEstate(c *fiber.Ctx) error {
	idString := c.Params("id")
	//get token
	accessToken := c.Cookies("accessToken")
	//get user id
	userID, err := helpers.ParseJWT(accessToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//
	err = estateDB.DeleteEstate(idString, userID)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted succesfully !",
	})

}
