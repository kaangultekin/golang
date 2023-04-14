package auth

import (
	"github.com/gofiber/fiber/v2"
	authService "golang/api/interfaces/auth"
	"strconv"
)

type AuthController struct {
	AuthService authService.IAuthService
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
