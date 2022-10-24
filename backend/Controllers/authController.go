// Package controllers Handles authentication of users and the validity of bearer tokens.
package controllers

import (
	databaseHandler "SmartWatch_Project/db"
	dbModels "SmartWatch_Project/db/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func hashAndSalt(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) error {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err
}
func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["username"] == "" || data["password"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Missing username or password!",
		})
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	user := &dbModels.User{
		Username: data["username"],
		Password: hashAndSalt(data["password"]),
	}

	err := databaseHandler.CreateUser(*user)
	if err != "" {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{"error": err})
	}

	return c.JSON(fiber.Map{"message": "Registered!"})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	user, err := dbModels.Users(dbModels.UserWhere.Username.EQ(data["username"])).OneG(context.Background())
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := comparePasswords(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"ID":       user.ID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}
