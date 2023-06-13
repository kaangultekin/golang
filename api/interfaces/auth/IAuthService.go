package auth

import (
	"github.com/gofiber/fiber/v2"
	authFormStructs "golang/api/structs/form/auth"
)

type IAuthService interface {
	Register(registerForm authFormStructs.RegisterFormStruct) (interface{}, error)
	Login(loginForm authFormStructs.LoginFormStruct) (interface{}, error)
	GetUser(id int) (interface{}, error)
	Logout(c *fiber.Ctx) (interface{}, error)
}
