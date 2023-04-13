package main

import (
	"github.com/gofiber/fiber/v2"
	"golang/routes"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
