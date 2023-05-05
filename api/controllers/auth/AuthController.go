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
