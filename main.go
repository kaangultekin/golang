package main

import (
	"github.com/gofiber/fiber/v2"
	"golang/api/config"
	"golang/api/routes"
	"golang/crons"
)

func init() {
	config.Connections()
}

func main() {
	app := fiber.New()

	crons.Cron()

	routes.Routes(app)

	app.Listen(":3000")
}
