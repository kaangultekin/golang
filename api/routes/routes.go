package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang/api/containers"
	authMiddlewares "golang/api/middlewares/auth"
	"golang/api/validations"
)

func Routes(app *fiber.App) {
	serviceContainer := containers.ServiceContainer()

	authMiddleware := authMiddlewares.AuthMiddleware()

	injectConnectedToAPIController := serviceContainer.InjectConnectedToAPIController()
	injectAuthController := serviceContainer.InjectAuthController()
	injectEndpointNotFoundController := serviceContainer.InjectEndpointNotFoundController()

	api := app.Group("/api")
	authControllerApi := api.Group("auth")

	/* Connected */
	app.All("/", injectConnectedToAPIController.ConnectedToAPI)
	/* Connected */

	/* Auth Controller */
	authControllerApi.Post("/register", validations.Validation("RegisterForm"), injectAuthController.Register)
	authControllerApi.Post("/login", validations.Validation("LoginForm"), injectAuthController.Login)
	authControllerApi.Get("/me", authMiddleware, injectAuthController.GetUser)
	/* Auth Controller */

	/* Not Found */
	app.Use(injectEndpointNotFoundController.EndpointNotFound)
	/* Not Found */
}
