package main

import (
	"github.com/gofiber/fiber/v2"
	"golang/api/config"
	"golang/api/migrations"
	"golang/api/routes"
)

func init() {
	config.ConnectDatabase()
	migrations.Migrate()
}

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
