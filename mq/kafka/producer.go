package kafka

import (
	"github.com/Shopify/sarama"
)

/*
ref:
	- github.com/Shopify/sarama
	- https://github.com/segmentio/kafka-go
	- https://github.com/confluentinc/confluent-kafka-go
	- github.com/bsm/sarama-cluster
		- 已废弃, 已合并到 sarama
		- https://github.com/Shopify/sarama/pull/1099
*/

type Producer struct {
	sync  sarama.SyncProducer  // 同步
	async sarama.AsyncProducer // 异步
}

func NewProducer() {

}

type ProducerGroup struct {
	Lite *ProducerLite
	Std  *ProducerStd
	Pro  *ProducerPro
}

// 低一致性:
type ProducerLite struct {
}

// 常规:
type ProducerStd struct {
}

// 高一致性:
type ProducerPro struct {
}
