package controllers

import (
	"SmartWatch_Project/ProtoBuffers"
	databaseHandler "SmartWatch_Project/db"
	dbModels "SmartWatch_Project/db/models"
	"encoding/hex"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type EventCBOR struct {
	ID             int    `json:"id"`
	Eventname      string `json:"eventname"`
	Eventtime      int64  `json:"eventtime"`
	CreatedAt      int64  `json:"created_at"`
	Userid         int    `json:"userid"`
	Issilent       bool   `json:"issilent"`
	Snoozedisabled bool   `json:"snoozedisabled"`
}

func GetEventsCBOR(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	err, eventsQ := databaseHandler.FetchMultipleByID(int(idString), c.Context())
	if err != "" {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	var cleanEvent []EventCBOR

	for _, j := range eventsQ {
		event := EventCBOR{
			ID:             j.ID,
			Eventname:      j.Eventname,
			Eventtime:      int64(j.Eventtime),
			Userid:         j.Userid,
			Issilent:       j.Issilent,
			Snoozedisabled: j.Snoozedisabled,
		}
		cleanEvent = append(cleanEvent, event)
	}

	marshal, err2 := cbor.Marshal(cleanEvent)
	if err2 != nil {
		return err2
	}

	return c.SendString(hex.EncodeToString(marshal))
}

func DeleteEventCBOR(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	decoded, err := hex.DecodeString(c.Params("bytes"))

	newEventProto := &ProtoBuffers.EventsOneEvent{}
	err = cbor.Unmarshal(decoded, newEventProto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed CBOR object!",
		})
	}

	errString := databaseHandler.DeleteEventByID(int(idString), int(newEventProto.Id), c.Context())
	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event Removed!"})

}

func DebugCBOR(c *fiber.Ctx) error {

	event := EventCBOR{
		ID:             123,
		Eventname:      "DebugEvent",
		Eventtime:      time.Now().Unix(),
		CreatedAt:      time.Now().Unix(),
		Userid:         1,
		Issilent:       true,
		Snoozedisabled: true,
	}

	marshal, err := cbor.Marshal(event)
	if err != nil {
		return err
	}

	var dat EventCBOR

	if err := cbor.Unmarshal(marshal, &dat); err != nil {
		panic(err)
	}

	return c.SendString(hex.EncodeToString(marshal))

}

func NewEventCBOR(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["ID"].(float64)
	fmt.Println(userID)

	decoded, err := hex.DecodeString(c.Params("bytes"))

	fmt.Println(hex.EncodeToString(decoded))

	var newEventCBOR EventCBOR
	err = cbor.Unmarshal(decoded, &newEventCBOR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed CBOR object!",
		})
	}

	eventError := databaseHandler.NewEventPopUp(c.Context(), dbModels.Event{
		ID:             newEventCBOR.ID,
		Eventname:      newEventCBOR.Eventname,
		Eventtime:      int(newEventCBOR.Eventtime),
		Userid:         newEventCBOR.Userid,
		Issilent:       newEventCBOR.Issilent,
		Snoozedisabled: newEventCBOR.Snoozedisabled,
	})

	if eventError != "" {
		c.Status(fiber.StatusConflict)
		return c.Status(400).JSON(fiber.Map{"error": eventError})
	}

	return c.JSON(fiber.Map{"message": "New event Added!"})

}

func ChangeEventTimeCBOR(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	decoded, err := hex.DecodeString(c.Params("bytes"))

	var newEventCBOR EventCBOR
	err = cbor.Unmarshal(decoded, &newEventCBOR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed CBOR object!",
		})
	}

	errString := databaseHandler.UpdateEventTime(int(idString), newEventCBOR.ID,
		time.Unix(newEventCBOR.Eventtime, 0), c.Context())

	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event time updated!"})

}
