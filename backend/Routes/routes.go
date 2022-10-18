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

	userRoute := api.Group("user")

	// JWT Middleware
	userRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	userRoute.Get("info", controllers.User)

	userRoute.Post("event", controllers.NewEvent)
	userRoute.Delete("event", controllers.DeleteEvent)
	userRoute.Patch("event", controllers.ChangeEventTime)

	userRoute.Get("events", controllers.GetEvents)

}
