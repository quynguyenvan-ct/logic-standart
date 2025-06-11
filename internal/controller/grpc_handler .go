package controller

import (
	"context"
	"golang/internal/entity"
	"golang/proto/kafka_pb"
)

type GRPCController struct {
	kafka_pb.UnimplementedKafkaServiceServer
	kafkaCtrl *KafkaController
}

func NewGRPCController(kafkaCtrl *KafkaController) *GRPCController {
	return &GRPCController{
		kafkaCtrl: kafkaCtrl,
	}
}

func (h *GRPCController) Publish(ctx context.Context, req *kafka_pb.PublishRequest) (response *kafka_pb.PublishResponse,err error) {
	event := &entity.KafkaEvent{
		TypeEvent: req.Event.TypeEvent,
		Payload:   req.Event.Payload,
	}

	if err := h.kafkaCtrl.Publish(ctx, req.Topic, event); err != nil {
		return &kafka_pb.PublishResponse{Message: "fail"}, err
	}

	return &kafka_pb.PublishResponse{Message: "success"}, nil
}
