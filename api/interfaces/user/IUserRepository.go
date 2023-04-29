package user

import (
	resultStructs "golang/api/structs/result"
)

type IUserRepository interface {
	GetByID(id int) (*resultStructs.ResultStruct, error)
}
