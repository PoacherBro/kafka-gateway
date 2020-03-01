package kafka

import "github.com/Shopify/sarama"

// AsyncProducer will push messages to kafka async
type AsyncProducer struct {
	ap sarama.AsyncProducer
}

// NewProducer create kafka async producer
func NewProducer(brokers []string) (*AsyncProducer, error) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		return nil, err
	}
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	return &AsyncProducer{
		ap: producer,
	}, nil
}

func (p *AsyncProducer) Send(msg *sarama.ProducerMessage) error {
	return nil
}
