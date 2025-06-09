package entity

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type KafkaEvent struct {
	TypeEvent string      `json:"TypeEvent"`
	Payload   interface{} `json:"payload"`
}