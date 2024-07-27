TAG := 1.0
IMAGE := agungsptr/go-asynq
CONTAINER := go-asynq
COMPOSE := docker compose -f docker-compose.yml

# Queue
build-queue:
	docker build -t $(IMAGE)_queue:$(TAG) queue

run-queue:
	@$(COMPOSE) down -v || true
	@$(COMPOSE) up -d --force-recreate queue

stop-queue:
	@docker container stop $(CONTAINER)_queue || true
	@docker container rm $(CONTAINER)_queue || true

# Services
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
