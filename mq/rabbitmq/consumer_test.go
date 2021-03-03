package rabbitmq

import (
	"fmt"
	"testing"
)

/*

web ui: http://localhost:15672/#/queues/exchange_server/test-queue
	- check web ui queue result

*/
func TestConsumer_Consume(t *testing.T) {

	c, err := NewConsumer(&ConnOption{
		Uri:      "amqp://rabbit:rabbit@localhost:5672/exchange_server",
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
		Vhost:    "",
	}, )

	t.Logf("producer create: %v", err)

	err = c.Consume(
		&Exchange{
			Name:       "test-exchange",
			Type:       "direct",
			Durable:    false,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		&Queue{
			Name:       "test-queue",
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false, // 是否具有排他性
			NoWait:     false, // 是否阻塞
			Args:       nil,   // 额外属性
		},
		"test-routing-key",
		"test-consumer-tag",
		func(message string) error {
			fmt.Printf("from taskFn call: %v", message)
			return nil
		},
	)

	t.Logf("consumer consume message: %v", err)

	//
	c.Close()

}
