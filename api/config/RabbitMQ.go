package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	messageConstants "golang/api/constants/message"
	"os"
)

var RMQ *amqp.Channel

func ConnectRabbitMQ() (bool, error) {
	envErr := godotenv.Load()

	if envErr != nil {
		return false, errors.New(messageConstants.ErrEnvFailed)
	}

	rabbitMQUsername := os.Getenv("RABBITMQ_USERNAME")
	rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")

	url := fmt.Sprintf("amqp://%s:%s@rabbitmq:%s/", rabbitMQUsername, rabbitMQPassword, rabbitMQPort)
	rabbitMQConn, rabbitMQErr := amqp.Dial(url)

	if rabbitMQErr != nil {
		rabbitMQConn.Close()
		return false, errors.New(rabbitMQErr.Error())
	}

	rabbitMQCh, rabbitMQChErr := rabbitMQConn.Channel()

	if rabbitMQChErr != nil {
		rabbitMQCh.Close()
		return false, errors.New(rabbitMQChErr.Error())
	}

	RMQ = rabbitMQCh

	return true, nil
}
