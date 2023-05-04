package auth

import (
	authFormStructs "golang/api/structs/form/auth"
	resultStructs "golang/api/structs/result"
)

type IAuthService interface {
	GetUserByID(id int) (*resultStructs.ResultStruct, error)
	Register(registerForm authFormStructs.RegisterFormStruct) (interface{}, error)
}
