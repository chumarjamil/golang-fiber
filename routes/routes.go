package routes

import (
	"github.com/gofiber/fiber/v2"
)

func setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", controller.Login)
	app.Get("/cashiers/:cashierId/logout", controller.Logut)
	app.Post("/cashiers/:cashierId/passcode", controller.Passcode)
}