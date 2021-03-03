package rabbitmq

import (
	"fmt"

	"github.com/better-go/pkg/log"
	"github.com/streadway/amqp"
)

//
//
//
type Producer struct {
	conn *amqp.Connection
}

//
//
//
func NewProducer(opt *ProducerOption) (*Producer, error) {
	p := new(Producer)

	connection, err := amqp.Dial(opt.MQUri())
	if err != nil {
		return p, fmt.Errorf("rabbitmq dial: %s", err)
	}

	p.conn = connection

	return p, nil

}

func (m *Producer) Publish(exchange *Exchange, queue *Queue, tag string, routingKey string, message string, reliable bool) error {
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
	if _, err := channel.QueueDeclare(queue.Name, queue.Durable, queue.AutoDelete, queue.Exclusive, queue.NoWait, queue.Args); err != nil {
		return fmt.Errorf("rabbitmq queue declare error: %v", err)
	}

	// 3. binding queue to exchange:
	if err := channel.QueueBind(queue.Name, routingKey, exchange.Name, false, nil, ); err != nil {
		return fmt.Errorf("rabbitmq queue binding exchange error: %v", err)
	}

	// 4. check:
	if reliable {
		log.Infof("enabling publishing confirms.")

		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("rabbitmq channel could not be put into confirm mode: %s", err)
		}

		//
		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		defer m.confirmOne(confirms)
	}

	// 5. do publish:
	log.Infof("rabbitmq do publish: msg=%v", message)
	if err := channel.Publish(
		exchange.Name,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    amqp.Persistent,
			Priority:        0,
			Body:            []byte(message),
		},
	); err != nil {
		return fmt.Errorf("rabbitmq exchange publish error: %v", err)
	}

	return nil
}

// One would typically keep a channel of publishings, a sequence number, and a
// set of unacknowledged sequence numbers and loop until the publishing channel
// is closed.
func (m *Producer) confirmOne(confirms <-chan amqp.Confirmation) {
	log.Infof("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		log.Infof("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Errorf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}

func (m *Producer) Close() {
	if m.conn != nil {
		defer m.conn.Close()
	}
}
