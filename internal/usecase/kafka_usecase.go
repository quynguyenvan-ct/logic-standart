package usecase

import (
	"context"
	"golang/internal/entity"
	"golang/internal/repository"
)

type KafkaProducer interface {
    Publish(ctx context.Context, topic string, event *entity.KafkaEvent) error
}

type KafkaUC interface {
	PublishUC(ctx context.Context, topic string, event *entity.KafkaEvent) error
}

type kafkaProducer struct {
	kafka KafkaProducer
}

func NewKafkaUC() KafkaUC{
	kafka := repository.NewKafkaProducer() 

	return &kafkaProducer {
		kafka:  kafka,
	}
}

func (kp *kafkaProducer) PublishUC(ctx context.Context, topic string, event *entity.KafkaEvent) error {
	if err := kp.kafka.Publish(ctx,topic,event); err != nil {
		return err
	}
	return nil
}