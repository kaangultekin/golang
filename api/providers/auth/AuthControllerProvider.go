package providers

import (
	authControllers "golang/api/controllers/auth"
	authInterfaces "golang/api/interfaces/auth"
)

func AuthControllerProvider(authService authInterfaces.IAuthService) *authControllers.AuthController {
	return &authControllers.AuthController{AuthService: authService}
}
