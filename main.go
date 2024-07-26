package main

import (
	"fmt"
	"log"

	"github.com/agungsptr/go-asynq/config"
	"github.com/agungsptr/go-asynq/tasks"
	"github.com/hibiken/asynq"
)

func main() {
	// Load .env file
	err := config.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	redisAddr := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, tasks.HandleImageDeliveryTask)
	mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
