package kafka

import (
	"github.com/Shopify/sarama"
)

type ProducerOptions struct {
	// meta
	Meta *sarama.Config

	// ext:
	key       string
	Group     string
	Topic     string
	Key       string
	Secret    string
	Batch     int64
	Cluster   string
	CreatedAt string
}

type ConsumerOptions struct {
}

type Meta struct {
	group       string
	topic       string
	mirrorTopic string
	cluster     string
	addr        string
	color       []byte
}
