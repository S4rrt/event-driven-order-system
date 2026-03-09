package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "orders",
	}
	return &Producer{
		writer: writer,
	}
}

func (p *Producer) Send(message []byte) error {
	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	},
	)
}
