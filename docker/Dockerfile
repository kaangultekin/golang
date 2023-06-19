FROM golang:1.19

RUN apt-get update && apt-get install -y npm redis-server
RUN npm install -g nodemon

RUN apt-get install -y apt-transport-https
RUN wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | apt-key add -
RUN echo "deb https://artifacts.elastic.co/packages/8.x/apt stable main" | tee -a /etc/apt/sources.list.d/elastic-8.x.list
RUN apt-get update && apt-get install -y elasticsearch logstash kibana

RUN echo "network.host: 0.0.0.0" >> /etc/elasticsearch/elasticsearch.yml

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

RUN go get -u github.com/gofiber/fiber/v2@v2.43.0

RUN go get github.com/joho/godotenv

RUN go get github.com/go-playground/validator/v10

RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres

RUN go get -u github.com/gofiber/jwt/v3
RUN go get -u github.com/golang-jwt/jwt/v4

COPY . .

EXPOSE 8090

ENV ENV=prod

CMD ["sh", "run.sh"]
