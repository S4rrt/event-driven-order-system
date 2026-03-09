package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer() *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "orders",
		GroupID:     "order-workers",
		StartOffset: kafka.FirstOffset,
	})
	return &Consumer{reader: reader}
}

func (c *Consumer) Read() ([]byte, error) {
	msg, err := c.reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	return msg.Value, nil
}
