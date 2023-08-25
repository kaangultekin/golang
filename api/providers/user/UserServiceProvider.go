package providers

import (
	userInterfaces "golang/api/interfaces/user"
	userServices "golang/api/services/user"
)

func UserServiceProvider(
	userRepository userInterfaces.IUserRepository,
	userRepositoryForElasticsearch userInterfaces.IUserRepositoryForElasticsearch,
) *userServices.UserService {
	return &userServices.UserService{
		UserRepository:                 userRepository,
		UserRepositoryForElasticsearch: userRepositoryForElasticsearch,
	}
}
