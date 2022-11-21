package controllers

import (
	"SmartWatch_Project/ProtoBuffers"
	databaseHandler "SmartWatch_Project/db"
	dbModels "SmartWatch_Project/db/models"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

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
			EventTimeUnix:  int64(j.Eventtime),
			UserId:         uint32(j.Userid),
			IsSilent:       j.Issilent,
			SnoozeDisabled: j.Snoozedisabled,
		}
		p.Events = append(p.Events, &event)
	}

	out, _ := proto.Marshal(&p)

	return c.SendString(hex.EncodeToString(out))
}

func DeleteEventProto(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	decoded, err := hex.DecodeString(c.Params("bytes"))

	newEventProto := &ProtoBuffers.EventsOneEvent{}
	err = proto.Unmarshal(decoded, newEventProto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed Protobuffer object!",
		})
	}

	errString := databaseHandler.DeleteEventByID(int(idString), int(newEventProto.Id), c.Context())
	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event Removed!"})

}

func Debug(c *fiber.Ctx) error {

	event := &ProtoBuffers.EventsOneEvent{
		Id:             uint32(6),
		EventName:      "j.Eventname",
		EventTimeUnix:  0,
		CreatedAtUnix:  0,
		UserId:         uint32(11),
		IsSilent:       true,
		SnoozeDisabled: true,
	}

	out, err := proto.Marshal(event)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	fmt.Printf("%v \n", out)

	newProtoBuff := &ProtoBuffers.EventsOneEvent{}
	err = proto.Unmarshal(out, newProtoBuff)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}

	return c.SendString(hex.EncodeToString(out))

}

func NewEventProto(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["ID"].(float64)
	fmt.Println(userID)

	decoded, err := hex.DecodeString(c.Params("bytes"))

	newEventProto := &ProtoBuffers.EventsOneEvent{}
	err = proto.Unmarshal(decoded, newEventProto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed Protobuffer object!",
		})
	}

	newEvent := dbModels.Event{
		Userid:         int(userID),
		Eventname:      newEventProto.EventName,
		Eventtime:      int(newEventProto.EventTimeUnix),
		Issilent:       newEventProto.IsSilent,
		Snoozedisabled: newEventProto.SnoozeDisabled,
	}

	eventError := databaseHandler.NewEventPopUp(c.Context(), newEvent)

	if eventError != "" {
		c.Status(fiber.StatusConflict)
		return c.Status(400).JSON(fiber.Map{"error": eventError})
	}

	return c.JSON(fiber.Map{"message": "New event Added!"})

}

func ChangeEventTimeProto(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idString := claims["ID"].(float64)

	decoded, err := hex.DecodeString(c.Params("bytes"))

	newEventProto := &ProtoBuffers.EventsOneEvent{}
	err = proto.Unmarshal(decoded, newEventProto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Malformed Protobuffer object!",
		})
	}

	errString := databaseHandler.UpdateEventTime(int(idString), int(newEventProto.Id),
		time.Unix(newEventProto.EventTimeUnix, 0), c.Context())

	if errString != "" {
		return c.Status(400).JSON(fiber.Map{"error": errString})
	}

	return c.JSON(fiber.Map{"message": "Event time updated!"})

}
