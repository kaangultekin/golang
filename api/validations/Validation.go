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
		registerForm       = &authFormStructs.RegisterFormStruct{}
		loginForm          = &authFormStructs.LoginFormStruct{}
		updateMeForm       = &authFormStructs.UpdateMeFormStruct{}
		updatePasswordForm = &authFormStructs.UpdatePasswordFormStruct{}
	)

	switch form {
	case "RegisterForm":
		return registerForm
	case "LoginForm":
		return loginForm
	case "UpdateMeForm":
		return updateMeForm
	case "UpdatePasswordForm":
		return updatePasswordForm
	default:
		return nil
	}
}
