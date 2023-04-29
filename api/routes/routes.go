package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang/api/containers"
	authValidations "golang/api/validations/auth"
)

func Routes(app *fiber.App) {
	serviceContainer := containers.ServiceContainer()

	injectConnectedToAPIController := serviceContainer.InjectConnectedToAPIController()
	injectAuthController := serviceContainer.InjectAuthController()
	injectEndpointNotFoundController := serviceContainer.InjectEndpointNotFoundController()

	api := app.Group("/api")
	authControllerApi := api.Group("auth")

	/* Connected */
	app.All("/", injectConnectedToAPIController.ConnectedToAPI)
	/* Connected */

	/* AuthController */
	authControllerApi.Get("/user/:id", injectAuthController.GetUserByID)
	authControllerApi.Post("/register", authValidations.RegisterFormValidation(), injectAuthController.Register)
	/* AuthController */

	/* Not Found */
	app.Use(injectEndpointNotFoundController.EndpointNotFound)
	/* Not Found */
}
