package controllers

import (
	"SmartWatch_Project/ProtoBuffers"
	databaseHandler "SmartWatch_Project/db"
	dbModels "SmartWatch_Project/db/models"
	"encoding/hex"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

func User(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)

	return c.SendString("Welcome " + name)

}

func GetEvents(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	err, eventsQ := databaseHandler.FetchMultipleByID(int(idString), c.Context())
	if err != "" {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	b, marshalErr := json.Marshal(eventsQ)

	if marshalErr != nil {
		return c.Status(400).JSON(fiber.Map{"error": marshalErr})
	}

	return c.JSON(string(b))

}

func GetEventsProto(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	err, eventsQ := databaseHandler.FetchMultipleByID(int(idString), c.Context())
	if err != "" {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	p := ProtoBuffers.Events{
		Events: []*ProtoBuffers.EventsOneEvent{},
	}

	for _, j := range eventsQ {
		event := ProtoBuffers.EventsOneEvent{
			Id:             uint32(j.ID),
			EventName:      j.Eventname,
			Eventtime:      timestamppb.New(j.Eventtime),
			CreatedAt:      timestamppb.New(j.CreatedAt.Time),
			UserId:         uint32(j.Userid),
			IsSilent:       j.Issilent,
			SnoozeDisabled: j.Snoozedisabled,
		}
		p.Events = append(p.Events, &event)
	}

	out, _ := proto.Marshal(&p)

	return c.JSON(hex.EncodeToString(out))

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

	errString := databaseHandler.UpdateEventTime(int(idString), int(eventID), eventUnixTime, c.Context())
	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event time updated!"})

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
	eventUnixTime := time.Unix(i, 0)

	newEvent := dbModels.Event{
		Userid:         int(userID),
		Eventname:      data["eventName"],
		Eventtime:      eventUnixTime,
		Issilent:       alarmType,
		Snoozedisabled: snoozeType,
	}

	eventError := databaseHandler.CreateEvent(c.Context(), newEvent)

	if eventError != "" {
		c.Status(fiber.StatusConflict)
		return c.Status(400).JSON(fiber.Map{"error": eventError})
	}

	return c.JSON(fiber.Map{"message": "New event Added!"})

}
