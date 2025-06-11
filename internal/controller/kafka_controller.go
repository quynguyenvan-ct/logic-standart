package controller

import (
	"context"
	"golang/internal/entity"
	"golang/internal/usecase"
)



type KafkaController struct {
	kafkaUC usecase.KafkaUC
}

func NewKafkaController(kafka usecase.KafkaUC) *KafkaController {
	return &KafkaController{kafkaUC: kafka}
}

func (kc *KafkaController) Publish(ctx context.Context,topic string, event *entity.KafkaEvent) error {
	if err := kc.kafkaUC.PublishUC(ctx,topic,event); err != nil {
		return err
	}
	return nil
}




