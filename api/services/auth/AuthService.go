package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang/api/config"
	"golang/api/constants/general"
	messageConstants "golang/api/constants/message"
	"golang/api/helpers"
	userInterfaces "golang/api/interfaces/user"
	userModels "golang/api/models/user"
	authFormStructs "golang/api/structs/form/auth"
	"time"
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

	userModelJson, err := json.Marshal(userModel)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	_, esErr := config.ES.Index("users_index", bytes.NewReader(userModelJson))

	if esErr != nil {
		return nil, errors.New(err.Error())
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

func (as *AuthService) Logout(c *fiber.Ctx) (interface{}, error) {
	token := helpers.GetToken(c)
	claims := helpers.ParseToken(c)

	exp := claims["exp"].(float64)
	expTime := time.Unix(int64(exp), 0)
	duration := expTime.Sub(time.Now())

	err := config.Redis.Set(c.Context(), token, token, duration).Err()

	if err != nil {
		return nil, errors.New(messageConstants.ErrFailedLogout)
	}

	return messageConstants.SuccessGeneralMessage, nil
}

func (as *AuthService) UpdateMe(updateMeForm authFormStructs.UpdateMeFormStruct, c *fiber.Ctx) (interface{}, error) {
	userId := helpers.GetUserId(c)
	checkEmail, _ := as.UserRepository.GetByEmail(updateMeForm.Email)

	if checkEmail != nil && userId != int(checkEmail.ID) {
		return nil, errors.New(messageConstants.ErrUsedEmail)
	}

	user, userErr := as.UserRepository.GetById(userId)

	if userErr != nil {
		return nil, userErr
	}

	user.Name = updateMeForm.Name
	user.Surname = updateMeForm.Surname
	user.Email = updateMeForm.Email

	updateErr := config.DB.Save(user)

	if updateErr.Error != nil {
		return nil, updateErr.Error
	}

	return user, nil
}

func (as *AuthService) UpdatePassword(updatePasswordForm authFormStructs.UpdatePasswordFormStruct, c *fiber.Ctx) (interface{}, error) {
	userId := helpers.GetUserId(c)
	user, err := as.UserRepository.GetById(userId)

	if err != nil {
		return nil, err
	}

	compareHashAndPassword, _ := helpers.CompareHashAndPassword(user.Password, updatePasswordForm.OldPassword)

	if !compareHashAndPassword {
		return nil, errors.New(messageConstants.ErrCompareHashAndOldPassword)
	}

	user.Password = helpers.HashPassword(updatePasswordForm.Password)

	updateErr := config.DB.Save(user)

	if updateErr.Error != nil {
		return nil, updateErr.Error
	}

	return true, nil
}

func (as *AuthService) GetUsers(c *fiber.Ctx) (interface{}, error) {
	var users []userModels.User

	getUsersForRedis, errForRedis := config.Redis.Get(c.Context(), "users").Result()

	if errForRedis == nil {
		if errUnmarshalUsers := json.Unmarshal([]byte(getUsersForRedis), &users); errUnmarshalUsers != nil {
			return nil, errUnmarshalUsers
		}

		return users, nil
	}

	getUsers, err := as.UserRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	marshalUsers, marshalUsersErr := json.Marshal(getUsers)

	if marshalUsersErr != nil {
		return nil, marshalUsersErr
	}

	setUsersErr := config.Redis.Set(c.Context(), "users", marshalUsers, general.OneDay).Err()

	if setUsersErr != nil {
		return nil, setUsersErr
	}

	return getUsers, nil
}
