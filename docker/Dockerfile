FROM golang:1.19

RUN apt-get update && apt-get install -y npm
RUN npm install -g nodemon

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app

EXPOSE 3000

ENV ENV=prod

CMD ["sh", "run.sh"]
