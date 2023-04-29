package migrations

import (
	"golang/api/config"
	userModels "golang/api/models/user"
)

func Migrate() {
	config.DB.AutoMigrate(
		userModels.User{},
	)
}
