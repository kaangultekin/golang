package user

import (
	"golang/api/structs"
)

type IUserRepository interface {
	GetByID(id int) (*structs.Result, error)
}
