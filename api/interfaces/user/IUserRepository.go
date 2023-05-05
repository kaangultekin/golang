package user

import (
	userModels "golang/api/models/user"
)

type IUserRepository interface {
	GetByEmail(email string) (*userModels.User, error)
}
