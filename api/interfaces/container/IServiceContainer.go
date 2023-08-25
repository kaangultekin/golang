package container

import (
	authControllers "golang/api/controllers/auth"
	generalControllers "golang/api/controllers/general"
	userControllers "golang/api/controllers/user"
)

type IServiceContainer interface {
	InjectUserController() *userControllers.UserController
	InjectAuthController() *authControllers.AuthController
	InjectConnectedToAPIController() *generalControllers.ConnectedToAPIController
	InjectEndpointNotFoundController() *generalControllers.EndpointNotFoundController
}
