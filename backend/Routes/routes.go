package routes

import (
	"SmartWatch_Project/Controllers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)
	api.Get("time", controllers.CurrentTime)

	userRoute := api.Group("user")
	protoBuff := api.Group("proto")
	CBOR := api.Group("cbor")

	// JWT Middleware
	userRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	protoBuff.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	CBOR.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	userRoute.Get("events", controllers.GetEvents)
	userRoute.Post("event", controllers.NewEvent)
	userRoute.Delete("event", controllers.DeleteEvent)
	userRoute.Delete("eventDate", controllers.DeleteEventDate)
	userRoute.Patch("eventTime", controllers.ChangeEventTime)
	userRoute.Patch("event", controllers.ChangeEvent)

	protoBuff.Get("events", controllers.GetEventsProto)
	protoBuff.Post("event/:bytes", controllers.NewEventProto)
	protoBuff.Patch("event/:bytes", controllers.ChangeEventTimeProto)
	protoBuff.Delete("event/:bytes", controllers.DeleteEventProto)

	protoBuff.Get("debug", controllers.Debug)

	CBOR.Get("events", controllers.GetEventsCBOR)
	CBOR.Post("event/:bytes", controllers.NewEventCBOR)
	CBOR.Patch("event/:bytes", controllers.ChangeEventTimeCBOR)
	CBOR.Delete("event/:bytes", controllers.DeleteEventCBOR)
	CBOR.Get("debug", controllers.DebugCBOR)
}
