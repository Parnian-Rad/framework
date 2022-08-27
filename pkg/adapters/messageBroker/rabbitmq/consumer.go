package rabbitmq

import (
	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
	"github.com/wagslane/go-rabbitmq"
)

type RabbitmqMessageConsumer struct {
	dsn  string
	conn *rabbitmq.Consumer
}

func NewConsumer(dsn string) ports.MessageConsumer {
	consumer, err := rabbitmq.NewConsumer(
		dsn, rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		panic(err)
	}

	return &RabbitmqMessageConsumer{
		dsn:  dsn,
		conn: &consumer,
	}
}

func (rmc *RabbitmqMessageConsumer) Consume(queue string, routeKeys []string, resultChan chan []byte) error {
	return rmc.conn.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			resultChan <- d.Body
			return rabbitmq.Ack
		},
		queue,
		routeKeys,
	)
}
