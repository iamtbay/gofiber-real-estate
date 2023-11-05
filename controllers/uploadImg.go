package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFile(c *fiber.Ctx) error {
	var imagesSlice []string
	//parse input
	files, err := c.MultipartForm()
	if err != nil {
		return err
	}

	for _, fieldHeader := range files.File {
		for _, file := range fieldHeader {
			//get extension
			extension := strings.Split(file.Filename, ".")
			//create file
			id := uuid.New()
			newFileName := fmt.Sprintf("%v.%v", id, extension[1])
			c.SaveFile(file, fmt.Sprintf("public/uploads/%s", newFileName))
			url := c.BaseURL()
			imagesSlice = append(imagesSlice, fmt.Sprintf("%v/images/%v", url, newFileName))

		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"images": imagesSlice})
}

