package providers

import (
	userRepository "golang/api/repositories/user"
)

func UserRepositoryProvider() *userRepository.UserRepository {
	return &userRepository.UserRepository{}
}
