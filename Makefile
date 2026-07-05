.PHONY: tidy run test cover docker up down logs

tidy:
	go mod tidy

run:
	go run ./cmd/api

test:
	go test ./...

cover:
	go test ./... -cover

docker:
	docker build -t ligafit-api .

up:
	docker compose up --build

down:
	docker compose down -v

logs:
	docker compose logs -f api