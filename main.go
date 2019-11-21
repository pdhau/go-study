package main

import (
	"context"
	"log"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})

	if err != nil {
		log.Fatalf("%v", err)
	}

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic name",
		SubscriptionName: "sub-1",
		Type:             pulsar.Shared,
	})

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer consumer.Close()

	ctx := context.Background()

	for {
		msg, err := consumer.Receive(ctx)
		if err != nil {
			log.Fatalf("%v", err)
		}

		err = ProcessMessage(msg)
		if err != nil {
			consumer.Ack(msg)
		} else {
			consumer.Nack(msg)
		}
	}
}

func ProcessMessage(msg pulsar.Message) error {

	return nil
}
