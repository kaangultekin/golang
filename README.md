WELCOME

Project Install

Step 1: Create .env

Step 2: sh run.sh

NOTES

Elasticsearch

* docker exec -it {ELASTICSEARCH_CONTAINER_ID} /bin/bash
* bin/elasticsearch-reset-password --username kibana_system -i 
* bin/elasticsearch-reset-password -u elastic

RabbitMQ

* RABBITMQ_DEFAULT_USER=guest
* RABBITMQ_DEFAULT_PASS=guest