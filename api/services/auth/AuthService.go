package auth

import (
	userRepository "golang/api/interfaces/user"
	userModel "golang/api/models/user"
)

type AuthService struct {
	UserRepository userRepository.IUserRepository
}

func (as *AuthService) GetUserByID(id int) (*userModel.User, error) {
	user, err := as.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
