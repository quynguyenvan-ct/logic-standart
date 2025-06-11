package entity

type KafkaEvent struct {
	TypeEvent []byte `json:"TypeEvent"`
	Payload   []byte `json:"payload"`
}