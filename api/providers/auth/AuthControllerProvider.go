package providers

import (
	authController "golang/api/controllers/auth"
	iAuthService "golang/api/interfaces/auth"
)

func AuthControllerProvider(authService iAuthService.IAuthService) *authController.AuthController {
	return &authController.AuthController{AuthService: authService}
}
