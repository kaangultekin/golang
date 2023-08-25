package user

import (
	"errors"
	"fmt"
	"golang/api/config"
	messageConstants "golang/api/constants/message"
	userModels "golang/api/models/user"
)

type UserRepository struct{}

func (ur *UserRepository) GetByEmail(email string) (*userModels.User, error) {
	var user userModels.User

	result := config.DB.Where("email", email).First(&user)

	if result.Error != nil {
		return nil, errors.New(messageConstants.ErrUserNotFound)
	}

	return &user, nil
}

func (ur *UserRepository) GetById(id int) (*userModels.User, error) {
	var user userModels.User

	result := config.DB.Where("id", id).
		Where("status", 1).
		First(&user)

	if result.Error != nil {
		return nil, errors.New(messageConstants.ErrUserNotFound)
	}

	return &user, nil
}

func (ur *UserRepository) SearchUsers(query string) (*[]userModels.User, error) {
	var users []userModels.User

	result := config.DB.Where(fmt.Sprintf("name ILIKE '%s'", "%"+query+"%")).
		Or(fmt.Sprintf("surname ILIKE '%s'", "%"+query+"%")).
		Or(fmt.Sprintf("email LIKE '%s'", "%"+query+"%")).
		Order("id DESC").
		Limit(100).
		Find(&users)

	if result.Error != nil {
		return nil, errors.New(messageConstants.ErrUserNotFound)
	}

	return &users, nil
}
