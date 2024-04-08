# Apache Kafka
## Project Installation
1.  One needs to install the _Confluent Go Kafka_ dependency from this project file root directory via
    ```bash
        go get github.com/confluentinc/confluent-kafka-go/kafka
    ```

## Kafka Setup
1.  The _Kafka broker_ can be started locally with
    ```bash
        docker-compose up -d
    ```
2.  All the settings for the Kafka bootstrap server can be configured inside the file `getting-started.properties`.

## Topic Creation
1.  Events in Kafka are organized and durably stored in named topics.
2.  A new topic, _purchases_ can be created using the `kafka-topics` command from the local running Kafka broker:
    ```bash
        docker compose exec broker \
            kafka-topics --create \
                --topic purchases \
                --bootstrap-server localhost:9092 \
                --replication-factor 1 \
                --partitions 1
    ```