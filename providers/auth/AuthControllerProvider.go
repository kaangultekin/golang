package providers

import (
	authController "golang/controllers/auth"
	iAuthService "golang/interfaces/auth"
)

func AuthControllerProvider(authService iAuthService.IAuthService) *authController.AuthController {
	return &authController.AuthController{AuthService: authService}
}
