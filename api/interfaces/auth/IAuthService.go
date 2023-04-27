package auth

import (
	"golang/api/structs/result"
)

type IAuthService interface {
	GetUserByID(id int) (*result.Result, error)
}
