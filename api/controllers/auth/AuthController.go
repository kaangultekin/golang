package auth

import (
	"github.com/gofiber/fiber/v2"
	authInterfaces "golang/api/interfaces/auth"
	authFormStructs "golang/api/structs/form/auth"
	resultStructs "golang/api/structs/result"
	"strconv"
)

type AuthController struct {
	AuthService authInterfaces.IAuthService
}

func (ac *AuthController) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	user, err := ac.AuthService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get user",
		})
	}

	return c.JSON(user)
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var (
		registerForm authFormStructs.RegisterFormStruct
		result       *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	c.BodyParser(&registerForm)

	register, err := ac.AuthService.Register(registerForm)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusBadRequest
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusCreated
	result.Message = "New user registered."
	result.Datas = register

	return c.Status(result.Code).JSON(result)
}
