package providers

import (
	userInterfaces "golang/api/interfaces/user"
	authServices "golang/api/services/auth"
)

func AuthServiceProvider(userRepository userInterfaces.IUserRepository) *authServices.AuthService {
	return &authServices.AuthService{UserRepository: userRepository}
}
