#!/bin/sh

docker-compose up --build
nodemon --exec go run main.go --signal SIGTERM