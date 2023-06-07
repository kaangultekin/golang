package main

import (
	"github.com/gofiber/fiber/v2"
	"golang/api/config"
	"golang/api/routes"
)

func init() {
	go config.ConnectPostgresDB()
}

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
