package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/real-estate-api/database"
	"github.com/iamtbay/real-estate-api/helpers"
	"github.com/iamtbay/real-estate-api/models"
	"golang.org/x/crypto/bcrypt"
)

var authDB = database.InitAuth()

// Login
func Login(c *fiber.Ctx) error {
	var userInfo *models.User
	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	//database operations
	dbUserInfo, err := authDB.Login(userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// create jwt cookie for user.
	tokenString := helpers.CreateJWT(dbUserInfo)
	// //create cookie with jwt
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int(time.Hour.Seconds() * 24 * 3),
		Expires:  time.Now().Add(time.Hour * 12),
		HTTPOnly: true,
	})
	//
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Welcome %v", userInfo.Email)})
}

// Register
func Register(c *fiber.Ctx) error {
	var userInfo *models.User
	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), 10)
	if err != nil {
		panic(err)
	}
	userInfo.Password = string(hashedPassword)
	//
	//database operations
	err = authDB.Register(userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Succesfully registered!"})
}

// Logout
func Logout(c *fiber.Ctx) error {
	//clear cookie
	c.Cookie(&fiber.Cookie{
		Name:    "accessToken",
		MaxAge:  -1,
		Expires: time.Now().Add(time.Second * -3),
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Succesffully logged out!"})
}
