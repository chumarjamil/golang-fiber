package main

import (
	"fmt"
	db "go_lang_rest_api/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Creating structure")

	db.Connect()

	app := fiber.New()
	app.Use(app)
}
