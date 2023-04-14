package providers

import (
	iUserRepository "golang/api/interfaces/user"
	authService "golang/api/services/auth"
)

func AuthServiceProvider(userRepository iUserRepository.IUserRepository) *authService.AuthService {
	return &authService.AuthService{UserRepository: userRepository}
}
