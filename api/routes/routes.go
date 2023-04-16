package routes

import (
	"github.com/gofiber/fiber/v2"
	authProviders "golang/api/providers/auth"
	userProviders "golang/api/providers/user"
)

func Routes(app *fiber.App) {

	userRepositoryProvider := userProviders.UserRepositoryProvider()
	authServiceProvider := authProviders.AuthServiceProvider(userRepositoryProvider)
	authControllerProvider := authProviders.AuthControllerProvider(authServiceProvider)

	app.Get("/user/:id", authControllerProvider.GetUserByID)
}
