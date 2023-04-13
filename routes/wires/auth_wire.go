//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	authController "golang/controllers/auth"
	iAuthService "golang/interfaces/auth"
	iUserRepository "golang/interfaces/user"
	authProviders "golang/providers/auth"
	userProviders "golang/providers/user"
	userRepository "golang/repositories/user"
	authService "golang/services/auth"
)

var authProviderSet = wire.NewSet(
	authProviders.AuthControllerProvider,
	authProviders.AuthServiceProvider,
	wire.Bind(new(iAuthService.IAuthService), new(*authService.AuthService)),
)

var userProviderSet = wire.NewSet(
	userProviders.UserRepositoryProvider,
	wire.Bind(new(iUserRepository.IUserRepository), new(*userRepository.UserRepository)),
)

func InitializeAuthController() *authController.AuthController {
	wire.Build(authProviderSet, userProviderSet)

	return &authController.AuthController{}
}
