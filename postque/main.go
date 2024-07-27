package main

import (
	"asynq-postque/config"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeEmailDelivery = "email:deliver"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

func NewEmailDeliveryTask(userID int, templateID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{
		UserID:     userID,
		TemplateID: templateID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func main() {
	// Load .env file
	config.Load()

	redisAddr := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})

	// Creating task
	task, err := NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	// Enqueue task to be processed immediately
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// Schedule task to be processed in the future
	info, err = client.Enqueue(task, asynq.ProcessIn(1*time.Minute), asynq.Queue("critical"))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
