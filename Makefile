.PHONY: help build run test clean lint lint-backend lint-frontend lint-fix lint-fix-backend lint-fix-frontend

# Переменные
DC=docker-compose
DC_FILE=deployments/docker-compose.yml
DC_DEV_FILE=deployments/docker-compose.dev.yml

help: ## Показать помощь
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Собрать все контейнеры
	$(DC) -f $(DC_FILE) build

dev: ## Запустить в режиме разработки
	$(DC) -f $(DC_FILE) -f $(DC_DEV_FILE) up -d

prod: ## Запустить в production режиме
	$(DC) -f $(DC_FILE) up -d

down: ## Остановить все контейнеры
	$(DC) -f $(DC_FILE) down

logs: ## Показать логи
	$(DC) -f $(DC_FILE) logs -f

test-backend: ## Запустить тесты бэкенда
	cd backend && go test -v ./...

test-frontend: ## Запустить тесты фронтенда
	cd frontend && npm test

lint-backend: ## Проверить код бэкенда
	cd backend && golangci-lint run ./...

lint-frontend: ## Проверить код фронтенда
	cd frontend && npm run lint

lint: lint-backend lint-frontend

lint-fix-backend:
	cd backend && golangci-lint run --fix ./...

lint-fix-frontend:
	cd frontend && npm run lint:fix

lint-fix: lint-fix-backend lint-fix-frontend

docs: ## Сгенерировать документацию
	cd backend && swag init -g cmd/app/main.go -o ../docs/api/swagger
	cd docs && mdbook build

clean: ## Очистить все собранные файлы и контейнеры
	$(DC) -f $(DC_FILE) down -v
	rm -rf backend/bin/*
	rm -rf frontend/build/* 