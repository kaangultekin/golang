package auth

import (
	userModel "golang/api/models/user"
)

type IAuthService interface {
	GetUserByID(id int) (*userModel.User, error)
}
