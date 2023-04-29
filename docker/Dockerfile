FROM golang:1.19

RUN apt-get update && apt-get install -y npm
RUN npm install -g nodemon

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

RUN go get -u github.com/gofiber/fiber/v2@v2.43.0

RUN go get github.com/joho/godotenv

RUN go get github.com/go-playground/validator/v10

RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres

COPY . .

EXPOSE 8090

ENV ENV=prod

CMD ["sh", "run.sh"]
