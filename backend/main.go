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

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://127.0.0.1:8000, http://127.0.0.1:8000,  http://localhost:8000,  http://localhost:3000,",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
