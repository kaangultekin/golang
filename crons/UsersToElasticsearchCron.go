package crons

import (
	"bytes"
	"fmt"
	"golang/api/config"
)

func UsersToElasticcearchCron() {
	users, usersErr := config.RMQ.Consume(
		"UsersQueue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if usersErr != nil {
		fmt.Println(usersErr.Error())
	}

	for user := range users {
		es, esErr := config.ES.Index("users_index", bytes.NewReader(user.Body))

		if es.StatusCode == 201 {
			user.Ack(true)

			continue
		}

		if esErr != nil {
			fmt.Println(esErr.Error())
		}
	}
}
