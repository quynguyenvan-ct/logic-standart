package repository

import (
	"context"
	"fmt"
	"golang/config"
	"golang/internal/entity"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
    Publish(ctx context.Context, topic string, event *entity.KafkaEvent) error
}

type KafkaConsumer interface {
    Start(ctx context.Context) error
    Stop(ctx context.Context) error
}


type kafkaProducer struct {
    writer *kafka.Writer
}


func NewKafkaProducer() KafkaProducer {
	var brokers []string
	brokers = config.LoadConfig().Kafka.Brokers
    w := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  brokers,
        Balancer: &kafka.LeastBytes{},
    })
    return &kafkaProducer{writer: w}
}

func (p *kafkaProducer) Publish(ctx context.Context, topic string, event *entity.KafkaEvent) error {
    msg := kafka.Message{
        Topic: topic,
        Key:   []byte(event.TypeEvent),
        Value: event.Payload,
    }
    fmt.Print(topic)
    return p.writer.WriteMessages(ctx, msg)
}