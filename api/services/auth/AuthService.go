package auth

import (
	"errors"
	"golang/api/config"
	messageConstants "golang/api/constants/message"
	"golang/api/helpers"
	userInterfaces "golang/api/interfaces/user"
	userModels "golang/api/models/user"
	authFormStructs "golang/api/structs/form/auth"
)

type AuthService struct {
	UserRepository userInterfaces.IUserRepository
}

func (as *AuthService) Register(registerForm authFormStructs.RegisterFormStruct) (interface{}, error) {
	user, _ := as.UserRepository.GetByEmail(registerForm.Email)

	if user != nil {
		return nil, errors.New(messageConstants.ErrUsedEmail)
	}

	var userModel = &userModels.User{
		Name:     registerForm.Name,
		Surname:  registerForm.Surname,
		Email:    registerForm.Email,
		Password: helpers.HashPassword(registerForm.Password),
	}

	createErr := config.DB.Create(userModel)

	if createErr.Error != nil {
		return nil, createErr.Error
	}

	return userModel, nil
}
