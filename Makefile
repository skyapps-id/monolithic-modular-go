APP_NAME := monolithic-modular-go

.PHONY: run build test lint tidy docker docker-inventory migrate migrate-status migrate-down migrate-drop swag

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

build-inventory:
	go build -o bin/inventory cmd/inventory/main.go

swag:
	swag init -g docs/docs.go -o docs --parseInternal

migrate:
	go run cmd/migrate/main.go -up

migrate-status:
	go run cmd/migrate/main.go -status

migrate-down:
	go run cmd/migrate/main.go -down

migrate-drop:
	go run cmd/migrate/main.go -drop

docker:
	@echo "Usage: make docker SERVICE=server|inventory"
	docker build --build-arg SERVICE=$(SERVICE) --target runtime -t $(APP_NAME)-$(SERVICE) .

docker-inventory:
	docker build --build-arg SERVICE=inventory --target runtime -t $(APP_NAME)-inventory .

test:
	go test ./...

lint:
	golangci-lint run

tidy:
	go mod tidy