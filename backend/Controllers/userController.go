package controllers

import (
	databaseHandler "SmartWatch_Project/db"
	dbModels "SmartWatch_Project/db/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

var UTC, _ = time.LoadLocation("UTC")

func CurrentTime(c *fiber.Ctx) error {
	timeNow := time.Now().Unix()
	return c.SendString(strconv.FormatInt(timeNow, 10))

}

func GetEvents(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	err, eventsQ := databaseHandler.FetchMultipleByID(int(idString), c.Context())
	if err != "" {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(eventsQ)

}

func ChangeEvent(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["ID"].(float64)
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	if data["eventName"] == "" || data["eventTimeOldUnix"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing event name or event time!",
		})
	}

	if data["isSilent"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing alarm type (silent/loud)!",
		})
	}

	if data["SnoozeDisabled"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing snooze info!",
		})
	}

	snoozeType, err := strconv.ParseBool(data["SnoozeDisabled"])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed snooze setting for a new event!",
		})
	}

	alarmType, err := strconv.ParseBool(data["isSilent"])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed alarm type for a new event!",
		})
	}

	i, err := strconv.ParseInt(data["eventTimeNewUnix"], 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed UnixTime for a new event!",
		})
	}

	if i > 2068690600000 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed UnixTime for a new event!",
		})
	}

	eventTimeNewUnix := time.Unix(i, 0).In(UTC)

	newEvent := dbModels.Event{
		Userid:         int(userID),
		Eventname:      data["eventName"],
		Eventtime:      int(eventTimeNewUnix.Unix()),
		Issilent:       alarmType,
		Snoozedisabled: snoozeType,
	}

	errString := databaseHandler.UpdateEvent(int(userID), data["eventTimeOldUnix"], newEvent, c.Context())
	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event time updated!"})

}

func ChangeEventTime(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	if data["eventID"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing eventID!",
		})
	}

	eventID, err := strconv.ParseInt(data["eventID"], 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed eventID for a new event!",
		})
	}

	i, err := strconv.ParseInt(data["eventTime"], 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed UnixTime for a new event!",
		})
	}

	eventUnixTime := time.Unix(i, 0)

	if eventUnixTime.Unix() > 2068725233 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed UnixTime for a new event!",
		})
	}

	errString := databaseHandler.UpdateEventTime(int(idString), int(eventID), eventUnixTime, c.Context())
	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event time updated!"})

}

func DeleteEventDate(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	if data["eventDate"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing eventDate!",
		})
	}

	errString := databaseHandler.DeleteEventByDate(int(idString), data["eventDate"], c.Context())
	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event Removed!"})

}

func DeleteEvent(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	if data["eventID"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing eventID!",
		})
	}

	eventID, err := strconv.ParseInt(data["eventID"], 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed eventID for a new event!",
		})
	}

	errString := databaseHandler.DeleteEventByID(int(idString), int(eventID), c.Context())
	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event Removed!"})

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
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing event name or event time!",
		})
	}

	if data["isSilent"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing alarm type (silent/loud)!",
		})
	}

	if data["SnoozeDisabled"] == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing snooze info!",
		})
	}

	snoozeType, err := strconv.ParseBool(data["SnoozeDisabled"])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed snooze setting for a new event!",
		})
	}

	alarmType, err := strconv.ParseBool(data["isSilent"])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed alarm type for a new event!",
		})
	}

	i, err := strconv.ParseInt(data["eventTime"], 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed UnixTime for a new event!",
		})
	}

	eventUnixTime := time.Unix(i, 0).In(UTC)

	if i > 2068690600000 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed UnixTime for a new event!",
		})
	}

	newEvent := dbModels.Event{
		Userid:         int(userID),
		Eventname:      data["eventName"],
		Eventtime:      int(eventUnixTime.Unix()),
		Issilent:       alarmType,
		Snoozedisabled: snoozeType,
	}

	eventError := databaseHandler.NewEventPopUp(c.Context(), newEvent)

	if eventError != "" {
		c.Status(fiber.StatusConflict)
		return c.Status(400).JSON(fiber.Map{"error": eventError})
	}

	return c.JSON(fiber.Map{"message": "New event Added!"})

}
