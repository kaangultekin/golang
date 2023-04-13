package user

import (
	UserModel "golang/models/user"
)

type UserRepository struct{}

func (ur *UserRepository) GetByID(id int) (*UserModel.User, error) {
	user := &UserModel.User{
		ID:       id,
		Name:     "Kaan",
		Surname:  "G",
		Password: "1234",
		Email:    "kaangultekin1907@gmail.com",
	}

	return user, nil
}
