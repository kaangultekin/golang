package main

import (
	"github.com/gofiber/fiber/v2"
	"golang/api/config"
	"golang/api/routes"
)

var postgresDb = make(chan bool)

func init() {
	go func() {
		config.ConnectPostgresDB()
		postgresDb <- true
	}()
}

func main() {
	if !<-postgresDb {
		return
	}

	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
