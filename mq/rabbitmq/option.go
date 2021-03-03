package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

/*
ref:
	- https://github.com/streadway/amqp
		- usage: https://github.com/streadway/amqp/blob/master/_examples/simple-consumer/consumer.go
	- https://github.com/devimteam/amqp
	- https://github.com/koding/rabbitmq
*/

// producer:
type ProducerOption struct {
	// meta
	Meta *ConnOption

	//
	Exchange *Exchange
	Queue    *Queue

	// The key that when publishing a message to a exchange/queue will be only delivered to
	// given routing key listeners
	RoutingKey string

	// Publishing tag
	Tag string

	// Queue should be on the server/broker
	Mandatory bool

	// Consumer should be bound to server
	Immediate bool

	// Do not wait for a consumer
	NoWait bool

	// App specific data
	Args amqp.Table
}


// mq uri:
func (m *ProducerOption) MQUri() string {
	if m.Meta.Uri != "" {
		return m.Meta.Uri
	}

	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/", m.Meta.Username, m.Meta.Password, m.Meta.Host, m.Meta.Port)
	if m.Meta.Vhost != "" {
		uri += m.Meta.Vhost
	}
	return uri
}

// consumer:
type ConsumerOption struct {
	// meta
	Conn *ConnOption
}

// metadata:
type ConnOption struct {
	Uri      string // AMQP URI
	Host     string //
	Port     int
	Username string
	Password string
	Vhost    string
}

//////////////////////////////////////////////////////////////////////////////////////////

type Exchange struct {
	// Exchange name
	Name string

	// Exchange type
	Type string

	// Durable exchanges will survive server restarts
	Durable bool

	// Will remain declared when there are no remaining bindings.
	AutoDelete bool

	// Exchanges declared as `internal` do not accept accept publishing.Internal
	// exchanges are useful for when you wish to implement inter-exchange topologies
	// that should not be exposed to users of the broker.
	Internal bool

	// When noWait is true, declare without waiting for a confirmation from the server.
	NoWait bool

	// amqp.Table of arguments that are specific to the server's implementation of
	// the exchange can be sent for exchange types that require extra parameters.
	Args amqp.Table
}

type Queue struct {
	// The queue name may be empty, in which the server will generate a unique name
	// which will be returned in the Name field of Queue struct.
	Name string

	// Check Exchange comments for durable
	Durable bool

	// Check Exchange comments for auto delete
	AutoDelete bool

	// Exclusive queues are only accessible by the connection that declares them and
	// will be deleted when the connection closes.  Channels on other connections
	// will receive an error when attempting declare, bind, consume, purge or delete a
	// queue with the same name.
	Exclusive bool

	// When noWait is true, the queue will assume to be declared on the server. A
	// channel exception will arrive if the conditions are met for existing queues
	// or attempting to modify an existing queue from a different connection.
	NoWait bool

	// Check Exchange comments for Args
	Args amqp.Table
}
