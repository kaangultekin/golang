package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang/api/containers"
	authMiddlewares "golang/api/middlewares/auth"
	"golang/api/validations"
)

func Routes(app *fiber.App) {
	serviceContainer := containers.ServiceContainer()

	injectConnectedToAPIController := serviceContainer.InjectConnectedToAPIController()
	injectAuthController := serviceContainer.InjectAuthController()
	injectEndpointNotFoundController := serviceContainer.InjectEndpointNotFoundController()

	api := app.Group("/api")
	usersApi := api.Group("users")
	authApi := api.Group("auth")

	app.All("/", injectConnectedToAPIController.ConnectedToAPI)

	usersApi.Post("/register", validations.Validation("RegisterForm"), injectAuthController.Register)
	usersApi.Post("/login", validations.Validation("LoginForm"), injectAuthController.Login)
	usersApi.Get("/", injectAuthController.GetUsers)

	authApi.Use(authMiddlewares.AuthMiddleware())
	{
		authApi.Get("/me", injectAuthController.Me)
		authApi.Post("/logout", injectAuthController.Logout)
		authApi.Put("/update-me", validations.Validation("UpdateMeForm"), injectAuthController.UpdateMe)
		authApi.Put("/update-password", validations.Validation("UpdatePasswordForm"), injectAuthController.UpdatePassword)
	}

	app.Use(injectEndpointNotFoundController.EndpointNotFound)
}
