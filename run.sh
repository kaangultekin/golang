#!/bin/sh

docker-compose up -d --build
nodemon --exec go run main.go --signal SIGTERM
