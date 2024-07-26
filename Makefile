TAG := 1.0
IMAGE := agungsptr/go-asynq
CONTAINER := go-asynq
COMPOSE := docker compose -f docker-compose.yml

build:
	docker build -t $(IMAGE):$(TAG) .

infra:
	@echo "Starting redis..."
	@docker compose down -v  || true
	@docker compose up -d --force-recreate redis
	@echo "Redis is up ✅"

compose-up:
	@echo "Starting services..."
	@TAG=$(TAG) $(COMPOSE) down -v || true
	@TAG=$(TAG) $(COMPOSE) up -d --force-recreate

compose-down:
	@TAG=$(TAG) $(COMPOSE) down -v || true

purge:
	@make -s compose-down
	@docker image rm $(IMAGE):$(TAG) || true