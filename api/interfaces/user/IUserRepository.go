package user

import (
	"golang/api/structs/result"
)

type IUserRepository interface {
	GetByID(id int) (*result.Result, error)
}
