package config

import (
	"fmt"
	generalConstants "golang/api/constants/general"
	"time"
)

func Connections() {
	maxRetry := 10

	postgresDB := make(chan bool)
	redisDB := make(chan bool)
	es := make(chan bool)

	var postgresDBErr error
	var redisDBErr error
	var esErr error

	go func() {
		for try := 1; try <= maxRetry; try++ {
			connectPostgresDB, postgresErr := ConnectPostgresDB()
			connectRedisDB, redisErr := ConnectRedisDB()
			connectElasticsearch, elasticsearchErr := ConnectElasticsearch()

			if connectPostgresDB && connectRedisDB && connectElasticsearch {
				RunMigrations()

				postgresDB <- true
				redisDB <- true
				es <- true

				break
			}

			if try == maxRetry {
				postgresDBErr = postgresErr
				redisDBErr = redisErr
				esErr = elasticsearchErr

				postgresDB <- false
				redisDB <- false
				es <- false
			} else {
				if postgresErr != nil {
					fmt.Println(postgresErr)
				}

				if redisErr != nil {
					fmt.Println(redisErr)
				}

				if elasticsearchErr != nil {
					fmt.Println(redisErr)
				}

				time.Sleep(generalConstants.FiveSeconds)
			}
		}
	}()

	if !<-postgresDB {
		panic(postgresDBErr)
	}

	if !<-redisDB {
		panic(redisDBErr)
	}

	if !<-es {
		panic(esErr)
	}

	close(postgresDB)
	close(redisDB)
	close(es)
}
