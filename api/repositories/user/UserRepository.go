package user

import (
	"golang/api/config"
	UserModel "golang/api/models/user"
	"golang/api/structs"
)

type UserRepository struct {
}

func (ur *UserRepository) GetByID(id int) (*structs.Result, error) {
	result := &structs.Result{
		Datas: &UserModel.User{},
	}

	config.DB.Table("users").
		Select("id, status, name, surname, email, created_at, updated_at").
		Where("id = ?", id).
		Order("id desc").
		First(&result.Datas)

	return result, nil
}
