package routes

import (
	"github.com/gofiber/fiber/v2"
	authProviders "golang/api/providers/auth"
	userProviders "golang/api/providers/user"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")

	userRepositoryProvider := userProviders.UserRepositoryProvider()
	authServiceProvider := authProviders.AuthServiceProvider(userRepositoryProvider)
	authControllerProvider := authProviders.AuthControllerProvider(authServiceProvider)

	api.Get("/user/:id", authControllerProvider.GetUserByID)
}
