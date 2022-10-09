package routes

import (
	"SmartWatch_Project/Controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	api.Post("register", controllers.Register)
	//api.Post("login", controllers.Login)

	/*	userRoute := api.Group("user")

		userRouteAuth := userRoute.Use(middlewares.IsAuth)
		userRouteAuth.Get("info", controllers.User)
		userRouteAuth.Post("logout", controllers.LogOut)
		userRouteAuth.Put("info", controllers.UpdateInfo)
		userRouteAuth.Put("password", controllers.UpdatePassword)
	*/
}
