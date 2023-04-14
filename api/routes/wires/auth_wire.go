//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	authController "golang/api/controllers/auth"
	iAuthService "golang/api/interfaces/auth"
	iUserRepository "golang/api/interfaces/user"
	authProviders "golang/api/providers/auth"
	userProviders "golang/api/providers/user"
	userRepository "golang/api/repositories/user"
	authService "golang/api/services/auth"
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
