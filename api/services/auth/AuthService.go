package auth

import (
	userRepository "golang/api/interfaces/user"
	"golang/api/structs"
)

type AuthService struct {
	UserRepository userRepository.IUserRepository
}

func (as *AuthService) GetUserByID(id int) (*structs.Result, error) {
	user, err := as.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
