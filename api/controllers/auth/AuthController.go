package auth

import (
	"github.com/gofiber/fiber/v2"
	messageConstants "golang/api/constants/message"
	authInterfaces "golang/api/interfaces/auth"
	authFormStructs "golang/api/structs/form/auth"
	resultStructs "golang/api/structs/result"
)

type AuthController struct {
	AuthService authInterfaces.IAuthService
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var (
		registerForm authFormStructs.RegisterFormStruct
		result       *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	c.BodyParser(&registerForm)

	user, err := ac.AuthService.Register(registerForm)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusBadRequest
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusCreated
	result.Message = messageConstants.SuccessNewUser
	result.Datas = user

	return c.Status(result.Code).JSON(result)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var (
		loginForm authFormStructs.LoginFormStruct
		result    *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	c.BodyParser(&loginForm)

	login, err := ac.AuthService.Login(loginForm)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusUnauthorized
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusOK
	result.Message = messageConstants.SuccessLogin
	result.Datas = login

	return c.Status(result.Code).JSON(result)
}

func (ac *AuthController) GetMe(c *fiber.Ctx) error {
	var (
		result *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	user, err := ac.AuthService.GetMe(c)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusNotFound
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusOK
	result.Message = messageConstants.SuccessGeneralMessage
	result.Datas = user

	return c.Status(result.Code).JSON(result)
}

func (ac *AuthController) UpdateMe(c *fiber.Ctx) error {
	var (
		updateMeForm authFormStructs.UpdateMeFormStruct
		result       *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	c.BodyParser(&updateMeForm)

	updateMe, err := ac.AuthService.UpdateMe(updateMeForm, c)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusBadRequest
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusOK
	result.Message = messageConstants.SuccessUpdateMe
	result.Datas = updateMe

	return c.Status(result.Code).JSON(result)
}

func (ac *AuthController) UpdatePassword(c *fiber.Ctx) error {
	var (
		updatePasswordForm authFormStructs.UpdatePasswordFormStruct
		result             *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	c.BodyParser(&updatePasswordForm)

	_, err := ac.AuthService.UpdatePassword(updatePasswordForm, c)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusBadRequest
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusOK
	result.Message = messageConstants.SuccessUpdatePassword

	return c.Status(result.Code).JSON(result)
}

func (ac *AuthController) Logout(c *fiber.Ctx) error {
	var (
		result *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	logout, err := ac.AuthService.Logout(c)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusBadRequest
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusOK
	result.Message = logout.(string)

	return c.Status(result.Code).JSON(result)
}
