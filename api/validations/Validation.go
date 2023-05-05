package validations

import (
	"github.com/gofiber/fiber/v2"
	messageConstants "golang/api/constants/message"
	authFormStructs "golang/api/structs/form/auth"
	resultStructs "golang/api/structs/result"
)

func Validation(form string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			result *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
		)

		validateForm := getValidateForm(form)

		if validateForm == nil {
			result.Success = false
			result.Code = fiber.StatusNotFound
			result.Message = messageConstants.ErrFormNotFound

			return c.Status(result.Code).JSON(result)
		}

		c.BodyParser(validateForm)

		validate := ValidateForm(validateForm)

		if !validate.Success {
			result.Success = validate.Success
			result.Code = fiber.StatusBadRequest
			result.Message = messageConstants.ErrValidationFailed
			result.Datas = validate.ErrorFields

			return c.Status(result.Code).JSON(result)
		}

		return c.Next()
	}
}

func getValidateForm(form string) interface{} {
	var (
		registerForm = &authFormStructs.RegisterFormStruct{}
	)

	switch form {
	case "RegisterForm":
		return registerForm
	default:
		return nil
	}
}
