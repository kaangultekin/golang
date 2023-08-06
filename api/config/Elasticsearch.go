package config

import (
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	messageConstants "golang/api/constants/message"
	"os"
)

var ES *elasticsearch.Client

func ConnectElasticsearch() (bool, error) {
	envErr := godotenv.Load()

	if envErr != nil {
		return false, errors.New(messageConstants.ErrEnvFailed)
	}

	elasticsearchProtocol := os.Getenv("ELASTICSEARCH_PROTOCOL")
	elasticsearchHost := os.Getenv("ELASTICSEARCH_HOST")
	elasticsearchPort := os.Getenv("ELASTICSEARCH_PORT")
	elasticsearchUsername := os.Getenv("ELASTICSEARCH_USERNAME")
	elasticsearchPassword := os.Getenv("ELASTICSEARCH_PASSWORD")

	addr := fmt.Sprintf("%s://%s:%s", elasticsearchProtocol, elasticsearchHost, elasticsearchPort)

	cfg := elasticsearch.Config{
		Addresses: []string{addr},
		Username:  elasticsearchUsername,
		Password:  elasticsearchPassword,
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		return false, errors.New(err.Error())
	}

	ES = es

	return true, nil
}
