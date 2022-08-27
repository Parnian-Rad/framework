package ports

type MessageConsumer interface {
	Consume(queue string, routeKeys []string, resultChan chan []byte) error
}

type MessageProducer interface {
	Publish(payload interface{}, routeKeys []string) error
}
