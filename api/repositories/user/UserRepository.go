package user

import (
	"golang/api/config"
	userModels "golang/api/models/user"
	resultStructs "golang/api/structs/result"
)

type UserRepository struct{}

func (ur *UserRepository) GetByID(id int) (*resultStructs.ResultStruct, error) {
	result := &resultStructs.ResultStruct{
		Datas: &userModels.User{},
	}

	config.DB.Table("users").
		Select("id, status, name, surname, email, created_at, updated_at").
		Where("id = ?", id).
		Order("id desc").
		First(&result.Datas)

	return result, nil
}
