package main

import "github.com/gofiber/fiber/v2"

func main() {
	// create instance of fiber
	app := fiber.New()

	// create http handler
	app.Get("/testApi", func (ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success" : "true",
			"message" : "Here is my first API in Fiber",
		})
	})

	// listen on port
	app.Listen(":9090")
}
