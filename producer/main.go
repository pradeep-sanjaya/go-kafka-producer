package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddress := "localhost:9092"
	topic := "topic1"
	numMessages := 10000
	message := "Hello"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	startTime := time.Now()

	for i := 0; i < numMessages; i++ {
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Value: []byte(message),
		})

		if err != nil {
			log.Printf("send failed, %s", err)
			panic(err)
		}

		log.Printf("sent message %d\n", i+1)
	}

	log.Printf("sent %d messages to Kafka topic '%s'\n", numMessages, topic)

	diff := time.Now().Sub(startTime)
	log.Printf("time taken: %s", diff)
}
