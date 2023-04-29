package container

import (
	authControllers "golang/api/controllers/auth"
	generalControllers "golang/api/controllers/general"
)

type IServiceContainer interface {
	InjectAuthController() *authControllers.AuthController
	InjectConnectedToAPIController() *generalControllers.ConnectedToAPIController
	InjectEndpointNotFoundController() *generalControllers.EndpointNotFoundController
}
