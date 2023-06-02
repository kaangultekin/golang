package user

import (
	"errors"
	"golang/api/config"
	messageConstants "golang/api/constants/message"
	userModels "golang/api/models/user"
)

type UserRepository struct{}

func (ur *UserRepository) GetByEmail(email string) (*userModels.User, error) {
	var user userModels.User

	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, errors.New(messageConstants.ErrUserNotFound)
	}

	return &user, nil
}

func (ur *UserRepository) GetById(id int) (*userModels.User, error) {
	var user userModels.User

	result := config.DB.Where("id = ? AND status = ?", id, 1).First(&user)

	if result.Error != nil {
		return nil, errors.New(messageConstants.ErrUserNotFound)
	}

	return &user, nil
}
