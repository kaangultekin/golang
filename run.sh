#!/bin/sh

if [ ! -f "bin/main" ]; then
  docker-compose --build
fi

docker-compose up

nodemon --exec go run main.go --signal SIGTERM