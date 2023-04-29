package containers

import (
	authControllers "golang/api/controllers/auth"
	generalControllers "golang/api/controllers/general"
	containerInterfaces "golang/api/interfaces/container"
	authProviders "golang/api/providers/auth"
	generalProviders "golang/api/providers/general"
	userProviders "golang/api/providers/user"
	"sync"
)

type Kernel struct{}

func (k *Kernel) InjectConnectedToAPIController() *generalControllers.ConnectedToAPIController {
	connectedToAPIControllerProvider := generalProviders.ConnectedToAPIControllerProvider()

	return connectedToAPIControllerProvider
}

func (k *Kernel) InjectAuthController() *authControllers.AuthController {
	userRepositoryProvider := userProviders.UserRepositoryProvider()
	authServiceProvider := authProviders.AuthServiceProvider(userRepositoryProvider)
	authControllerProvider := authProviders.AuthControllerProvider(authServiceProvider)

	return authControllerProvider
}

func (k *Kernel) InjectEndpointNotFoundController() *generalControllers.EndpointNotFoundController {
	endpointNotFoundControllerProvider := generalProviders.EndpointNotFoundControllerProvider()

	return endpointNotFoundControllerProvider
}

func ServiceContainer() containerInterfaces.IServiceContainer {
	var (
		kernel *Kernel
		once   sync.Once
	)

	if kernel == nil {
		once.Do(func() {
			kernel = &Kernel{}
		})
	}

	return kernel
}
