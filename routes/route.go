package routes

import (
	"github.com/blogbackend/wz/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//app.Use(middleware.IsAuthenticate)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
}
