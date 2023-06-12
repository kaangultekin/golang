package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	messageConstants "golang/api/constants/message"
	"os"
)

var Redis *redis.Client

func ConnectRedisDB() (bool, error) {
	envErr := godotenv.Load()

	if envErr != nil {
		return false, errors.New(messageConstants.ErrEnvFailed)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisMainPort := os.Getenv("REDIS_MAIN_PORT")

	addr := fmt.Sprintf("%s:%s", redisHost, redisMainPort)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()

	if err != nil {
		return false, errors.New(messageConstants.ErrConnectRedisFailed)
	}

	Redis = redisClient

	return true, nil
}
