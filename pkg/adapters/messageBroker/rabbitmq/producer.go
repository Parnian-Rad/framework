package rabbitmq

import (
	"encoding/json"

	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
	"github.com/wagslane/go-rabbitmq"
)

type RabbitmqMessageProducer struct {
	dsn  string
	conn *rabbitmq.Publisher
}

func NewProducer(dsn string) ports.MessageProducer {
	publisher, err := rabbitmq.NewPublisher(dsn, rabbitmq.Config{})
	if err != nil {
		panic(err)
	}

	return &RabbitmqMessageProducer{
		dsn:  dsn,
		conn: publisher,
	}
}

func (rmp *RabbitmqMessageProducer) Publish(payload interface{}, routeKeys []string) error {
	payloadB, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = rmp.conn.Publish(
		payloadB,
		routeKeys,
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsPersistentDelivery,
	)
	return err
}
