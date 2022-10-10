package controllers

import (
	databaseHandler "SmartWatch_Project/db"
	dbModels "SmartWatch_Project/db/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/volatiletech/null/v8"
	"strconv"
	"time"
)

func User(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)

	return c.SendString("Welcome " + name)

}

func NewEvent(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["ID"].(float64)
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	if data["eventName"] == "" || data["eventTime"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Missing event name or event time!",
		})
	}

	if data["isSilent"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Missing alarm type (silent/loud)!",
		})
	}

	alarmType, err := strconv.ParseBool(data["isSilent"])
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Malformed alarm type for a new event!",
		})
	}

	i, err := strconv.ParseInt(data["eventTime"], 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Malformed UnixTime for a new event!",
		})
	}
	eventUnixTime := time.Unix(i, 0)

	newEvent := dbModels.Event{
		Eventid:   int(userID),
		Eventname: data["eventName"],
		Eventtime: eventUnixTime,
		Issilent:  null.Bool{Bool: alarmType},
	}

	eventError := databaseHandler.CreateEvent(c.Context(), newEvent)

	if eventError != "" {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{"error": eventError})
	}

	return c.JSON(fiber.Map{"message": "New event Added!"})

}
