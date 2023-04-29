package providers

import (
	userRepositories "golang/api/repositories/user"
)

func UserRepositoryProvider() *userRepositories.UserRepository {
	return &userRepositories.UserRepository{}
}
