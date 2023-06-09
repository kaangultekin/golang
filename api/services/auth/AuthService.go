package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
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

func (as *AuthService) Login(loginForm authFormStructs.LoginFormStruct) (interface{}, error) {
	user, err := as.UserRepository.GetByEmail(loginForm.Email)

	if err != nil {
		return nil, err
	}

	compareHashAndPassword, _ := helpers.CompareHashAndPassword(user.Password, loginForm.Password)

	if !compareHashAndPassword {
		return nil, errors.New(messageConstants.ErrCompareHashAndPassword)
	}

	if user.Status == 0 {
		return nil, errors.New(messageConstants.ErrUserStatusPassive)
	}

	token, tokenErr := helpers.GenerateJWT(user)

	if tokenErr != nil {
		return nil, tokenErr
	}

	return fiber.Map{
		"user":  user,
		"token": token,
	}, nil
}

func (as *AuthService) GetUser(id int) (interface{}, error) {
	user, err := as.UserRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
