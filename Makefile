TAG := 1.0
IMAGE := agungsptr/go-asynq
CONTAINER := go-asynq
COMPOSE := docker compose -f docker-compose.yml

# Queue Service
build-queue:
	docker build -t $(IMAGE)_queue:$(TAG) queue

run-queue:
	@$(COMPOSE) up -d --force-recreate queue

stop-queue:
	@docker container stop $(CONTAINER)_queue || true
	@docker container rm $(CONTAINER)_queue || true


# Monitor Service
build-monitor:
	docker build -t $(IMAGE)_monitor:$(TAG) monitor

run-monitor:
	@$(COMPOSE) up -d --force-recreate monitor

stop-monitor:
	@docker container stop $(CONTAINER)_monitor || true
	@docker container rm $(CONTAINER)_monitor || true


# All services (Redis, Queue, Monitor)
run-services:
	@$(COMPOSE) down -v || true
	@$(COMPOSE) up -d --force-recreate

stop-services:
	@$(COMPOSE) down -v || true

purge-services:
	@make -s stop-services
	@docker image rm $(IMAGE)_queue:$(TAG) || true

# Others
redis-cli:
	@docker exec -it go-asynq_redis redis-cli
