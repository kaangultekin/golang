package auth

import (
	authFormStructs "golang/api/structs/form/auth"
)

type IAuthService interface {
	Register(registerForm authFormStructs.RegisterFormStruct) (interface{}, error)
}
