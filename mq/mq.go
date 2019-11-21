package mq

import "github.com/apache/pulsar/pulsar-client-go/pulsar"

type MQClient struct {
	client pulsar.Client
}

func NewMQClient() MQClient {
	client, _ := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	return MQClient{
		client: client,
	}
}

func (c *MQClient) RecieveMessage() {

}
