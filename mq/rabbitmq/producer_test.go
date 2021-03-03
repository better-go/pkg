package rabbitmq

import (
	"testing"
)

/*

web ui: http://localhost:15672/#/queues/exchange_server/test-queue
	- check web ui queue result

*/
func TestProducer(t *testing.T) {

	p, err := NewProducer(&ProducerOption{
		Meta: &ConnOption{
			Uri:      "amqp://rabbit:rabbit@localhost:5672/exchange_server",
			Host:     "",
			Port:     0,
			Username: "",
			Password: "",
			Vhost:    "",
		},
		Exchange: &Exchange{
			Name:       "",
			Type:       "",
			Durable:    false,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		Queue:      nil,
		RoutingKey: "",
		Tag:        "",
		Mandatory:  false,
		Immediate:  false,
		NoWait:     false,
		Args:       nil,
	})

	t.Logf("producer create: %v", err)

	err = p.Publish(
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
		"",
		"test-routing-key",
		"hello world, from producer publish",
		true,
	)

	t.Logf("producer publish message: %v", err)

	//
	p.Close()

}
