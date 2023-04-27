package auth

import (
	"golang/api/structs"
)

type IAuthService interface {
	GetUserByID(id int) (*structs.Result, error)
}
