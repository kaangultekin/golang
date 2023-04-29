package auth

import (
	"github.com/gofiber/fiber/v2"
	authInterfaces "golang/api/interfaces/auth"
	authFormStructs "golang/api/structs/form/auth"
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

	var registerForm authFormStructs.RegisterFormStruct
	c.BodyParser(&registerForm)

	return c.Status(fiber.StatusOK).JSON(registerForm)
}
