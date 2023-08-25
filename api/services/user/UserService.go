package user

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang/api/config"
	"golang/api/constants/general"
	userInterfaces "golang/api/interfaces/user"
	userModels "golang/api/models/user"
)

type UserService struct {
	UserRepository                 userInterfaces.IUserRepository
	UserRepositoryForElasticsearch userInterfaces.IUserRepositoryForElasticsearch
}

func (us *UserService) GetUser(id int) (interface{}, error) {
	user, err := us.UserRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) SearchUsers(query string, c *fiber.Ctx) (interface{}, error) {
	var (
		users    []userModels.User
		cacheKey = fmt.Sprintf("users:%s", query)
	)

	searchUsersForElasticsearch, searchUsersForElasticsearchErr := us.UserRepositoryForElasticsearch.SearchUsers(query)

	if searchUsersForElasticsearchErr == nil {
		return searchUsersForElasticsearch, nil
	}

	searchUsersForRedis, errForRedis := config.Redis.Get(c.Context(), cacheKey).Result()

	if errForRedis == nil {
		if errUnmarshalUsers := json.Unmarshal([]byte(searchUsersForRedis), &users); errUnmarshalUsers != nil {
			return nil, errUnmarshalUsers
		}

		return users, nil
	}

	searchUsers, err := us.UserRepository.SearchUsers(query)

	if err != nil {
		return nil, err
	}

	marshalUsers, marshalUsersErr := json.Marshal(searchUsers)

	if marshalUsersErr != nil {
		return nil, marshalUsersErr
	}

	setUsersErr := config.Redis.Set(c.Context(), cacheKey, marshalUsers, general.OneDay).Err()

	if setUsersErr != nil {
		return nil, setUsersErr
	}

	return searchUsers, nil
}
