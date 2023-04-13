package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang/routes/wires"
) //test

func Routes(app *fiber.App) {

	authController := wires.InitializeAuthController()

	app.Get("/user/:id", authController.GetUserByID)
}
