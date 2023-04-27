package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresName := os.Getenv("POSTGRES_DB")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresSslMode := os.Getenv("POSTGRES_SSL_MODE")
	postgresTimeZone := os.Getenv("POSTGRES_TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		postgresHost, postgresUser, postgresPassword, postgresName, postgresPort, postgresSslMode, postgresTimeZone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}
