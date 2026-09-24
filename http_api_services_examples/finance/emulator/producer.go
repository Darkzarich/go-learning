package main

import (
	"context"
	"fmt"
	"time"

	kafkaClient "github.com/segmentio/kafka-go"
)

type ProducerConfig struct {
	Brokers []string
	Topic   string
}

// Producer ...
type Producer interface {
	PushMessage(ctx context.Context, msgData ...[]byte) (err error)
}

type producer struct {
	kafka *kafkaClient.Writer
}

// PushMessage ...
func (p *producer) PushMessage(ctx context.Context, msgData ...[]byte) (err error) {
	if len(msgData) == 0 {
		return nil
	}

	now := time.Now()
	messages := make([]kafkaClient.Message, len(msgData))

	for i := range msgData {
		messages[i] = kafkaClient.Message{
			Time:  now,
			Value: msgData[i],
		}
	}

	if err = p.kafka.WriteMessages(ctx, messages...); err != nil {
		return fmt.Errorf("error write kafka message: %w", err)
	}

	return nil
}

// NewProducer ...
func NewProducer(cfg ProducerConfig) Producer {
	writer := &kafkaClient.Writer{
		Addr:                   kafkaClient.TCP(cfg.Brokers...),
		Topic:                  cfg.Topic,
		Balancer:               &kafkaClient.LeastBytes{},
		AllowAutoTopicCreation: true,
		BatchTimeout:           50 * time.Millisecond,
		RequiredAcks:           kafkaClient.RequireOne,
		Async:                  false,
	}

	return &producer{
		kafka: writer,
	}
}
