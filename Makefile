.PHONY: run build test docker up down

run:
	go run ./cmd/api

build:
	go build -o ligafit-api ./cmd/api

test:
	go test ./...

docker:
	docker build -t ligafit-api .

up:
	docker compose up --build

down:
	docker compose down -v