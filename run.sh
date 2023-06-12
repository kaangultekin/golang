#!/bin/sh

docker-compose build

docker-compose up

nodemon --exec go run main.go --signal SIGTERM