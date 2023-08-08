package helpers

import (
	"errors"
	"github.com/streadway/amqp"
	"golang/api/config"
)

func CreateQueue(
	name string,
	durable,
	autoDelete,
	exclusive,
	noWait bool,
	args amqp.Table,
	exchange string,
	mandatory,
	immediate bool,
	msg amqp.Publishing,
) (bool, error) {
	qd, err := config.RMQ.QueueDeclare(
		name,
		durable,
		autoDelete,
		exclusive,
		noWait,
		args,
	)

	if err != nil {
		return false, errors.New(err.Error())
	}

	err = config.RMQ.Publish(
		exchange,
		qd.Name,
		mandatory,
		immediate,
		msg,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
