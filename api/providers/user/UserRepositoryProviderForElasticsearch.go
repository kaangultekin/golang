package providers

import (
	userRepositories "golang/api/repositories/user"
)

func UserRepositoryProviderForElasticsearch() *userRepositories.UserRepositoryForElasticsearch {
	return &userRepositories.UserRepositoryForElasticsearch{}
}
