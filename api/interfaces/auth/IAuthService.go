package auth

import (
	resultStructs "golang/api/structs/result"
)

type IAuthService interface {
	GetUserByID(id int) (*resultStructs.ResultStruct, error)
}
