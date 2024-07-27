# go-asynq

This repo is an example of a Queue System with Monitoring and Client. The Queue Service builds with [asynq](https://github.com/hibiken/asynq) + monitoring using [asynqmon](https://github.com/hibiken/asynqmon). All service was integrated in docker compose.

## How To Use

1. Copy `.env.example` to `.env`, and you can configure the value like what you want.
2. Build `Queue Service` image

```sh
make build-queue
```

3. Build `Monitoring Service` image

```sh
make build-monitor
```

4. Now you can run all services (Redis, Queue, Monitor)

```sh
make run-services
```

5. To send queue task to system you can run this command

```sh
make run-client
```

## Explanation

### Queue Folder

This folder contains a queue system that is built using [asynq](https://github.com/hibiken/asynq). In this folder you can see examples of how to define a task with a handler and how to register it to the queue system.

### Queclient Folder

This folder contains how to create or add tasks to the queue system. There are 2 examples, first to enqueue tasks immediately and second with a time schedule.

### Monitor Folder

This folder contains monitoring service from [asynqmon](https://github.com/hibiken/asynqmon).

### Make Command Available
```sh
build-queue             # Build queue
run-queue               # Run queue only
stop-queue              # Stop queue only
build-monitor           # Build monitor
run-monitor             # Run monitor only
stop-monitor            # Stop monitor only
run-services            # Run all service (Redis, Queue, Monitor)
stop-services           # Stop all services
purge-services          # Purge all services
run-client              # Run client to create queue task
redis-cli               # Open redis cli
```
