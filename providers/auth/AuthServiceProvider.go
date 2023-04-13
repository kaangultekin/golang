package providers

import (
	iUserRepository "golang/interfaces/user"
	authService "golang/services/auth"
)

func AuthServiceProvider(userRepository iUserRepository.IUserRepository) *authService.AuthService {
	return &authService.AuthService{UserRepository: userRepository}
}
