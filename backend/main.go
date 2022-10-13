package main

import (
	routes "SmartWatch_Project/Routes"
	databaseHandler "SmartWatch_Project/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	databaseHandler.ConnectToDB()
	app := fiber.New()

	app.Static("/", "./public")

	corsSettings := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
		ExposeHeaders:    "Origin",
		AllowCredentials: true,
	})

	app.Use(corsSettings)

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
