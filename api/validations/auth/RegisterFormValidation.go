package auth

import (
	"github.com/gofiber/fiber/v2"
	authFormStructs "golang/api/structs/form/auth"
	"golang/api/validations"
)

func RegisterFormValidation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var registerForm authFormStructs.RegisterFormStruct
		c.BodyParser(&registerForm)

		result := validations.ValidateForm(registerForm)

		if !result.Success {
			return c.Status(fiber.StatusBadRequest).JSON(result)
		}

		return c.Next()
	}
}
