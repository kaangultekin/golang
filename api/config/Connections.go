package config

import (
	"fmt"
	"time"
)

func Connections() {
	maxRetry := 10

	postgresDB := make(chan bool)
	redisDB := make(chan bool)

	var postgresDBErr error
	var redisDBErr error

	go func() {
		for try := 1; try <= maxRetry; try++ {
			connectPostgresDB, postgresErr := ConnectPostgresDB()
			connectRedisDB, redisErr := ConnectRedisDB()

			if connectPostgresDB && connectRedisDB {
				RunMigrations()

				postgresDB <- true
				redisDB <- true

				break
			}

			if try == maxRetry {
				postgresDBErr = postgresErr
				redisDBErr = redisErr

				postgresDB <- false
				redisDB <- false
			} else {
				if postgresErr != nil {
					fmt.Println(postgresErr)
				}

				if redisErr != nil {
					fmt.Println(redisErr)
				}

				time.Sleep(time.Second * 5)
			}
		}
	}()

	if !<-postgresDB {
		panic(postgresDBErr)
	}

	if !<-redisDB {
		panic(redisDBErr)
	}

	close(postgresDB)
	close(redisDB)
}
