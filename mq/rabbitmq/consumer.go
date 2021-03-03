package rabbitmq

import (
	"fmt"
	"github.com/better-go/pkg/log"
	"github.com/streadway/amqp"
)

// task handler:
type TaskFunc func(message string) error

//
//
//
type Consumer struct {
	conn *amqp.Connection

	//
	done chan error
}

//
//
//
func NewConsumer(opt *ConnOption) (*Consumer, error) {
	p := new(Consumer)

	connection, err := amqp.Dial(opt.ConnUri())
	if err != nil {
		return p, fmt.Errorf("rabbitmq dial: %s", err)
	}

	p.conn = connection
	p.done = make(chan error)
	return p, nil

}

//
func (m *Consumer) Consume(exchange *Exchange, queue *Queue, routingKey string, consumerTag string, taskFn TaskFunc) error {
	channel, err := m.conn.Channel()
	if err != nil {
		return fmt.Errorf("rabbitmq get channel error: %v'", err)
	}
	defer channel.Close()

	// 1. declare exchange:
	if err := channel.ExchangeDeclare(
		exchange.Name,
		exchange.Type,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("rabbitmq exchange declare error: %v", err)
	}

	// 2. declare queue:
	if _, err := channel.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		queue.Args,
	); err != nil {
		return fmt.Errorf("rabbitmq queue declare error: %v", err)
	}

	// 3. binding queue to exchange:
	if err := channel.QueueBind(queue.Name, routingKey, exchange.Name, false, nil, ); err != nil {
		return fmt.Errorf("rabbitmq queue binding exchange error: %v", err)
	}

	// 4. consume:
	deliveries, err := channel.Consume(
		queue.Name,  // name
		consumerTag, // consumerTag,
		true,        // 需要打开, 不然会重复消费
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return fmt.Errorf("rabbitmq queue consume error: %s", err)
	}

	// 5. do handle task:
	go m.handleTask(deliveries, m.done, taskFn)
	return nil
}

// handle one task:
func (m *Consumer) handleTask(deliveries <-chan amqp.Delivery, done chan error, taskFn TaskFunc) {
	for d := range deliveries {
		log.Infof("rabbitmq consumer: handle task - size=%d B delivery=[%v] msg=%q", len(d.Body), d.DeliveryTag, d.Body)

		// handle one message:
		if err := taskFn(string(d.Body)); err != nil {
			log.Errorf("rabbitmq consumer: taskFunc error: %v", err)
		}

		// ack one:
		_ = d.Ack(false)
	}
	log.Infof("rabbitmq consumer handle done: deliveries channel closed")
	done <- nil
}

func (m *Consumer) Close() {
	if m.conn != nil {
		defer m.conn.Close()
	}
}
