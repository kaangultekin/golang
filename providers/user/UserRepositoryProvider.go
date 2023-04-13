package providers

import (
	userRepository "golang/repositories/user"
)

func UserRepositoryProvider() *userRepository.UserRepository {
	return &userRepository.UserRepository{}
}
