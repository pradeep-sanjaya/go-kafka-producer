package main

import (
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	brokerAddress := "localhost:9092,localhost:9093,localhost:9094"
	topic := "topic1"
	numMessages := 10000
	message := "Hello"

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokerAddress,
		//"key.serializer":    "org.apache.kafka.common.serialization.StringSerializer",
		//"value.serializer": "org.apache.kafka.common.serialization.StringSerializer",
		"acks":       "0",
		"batch.size": "100",
		//"linger.ms":  "100",
	})

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	deliveryChan := make(chan kafka.Event)

	startTime := time.Now()

	for i := 0; i < numMessages; i++ {
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, deliveryChan)

		if err != nil {
			return
		}

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			log.Printf(
				"Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic,
				m.TopicPartition.Partition,
				m.TopicPartition.Offset)
		}

		log.Printf("sent message %d\n", i+1)
	}

	log.Printf("sent %d messages to Kafka topic '%s'\n", numMessages, topic)

	diff := time.Now().Sub(startTime)
	log.Printf("time taken: %s", diff)
}
