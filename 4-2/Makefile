# Makefile

.PHONY: run lint test migrate-up migrate-down install

install:
	brew install golang-migrate

run:
	go run ./cmd/app/main.go

lint:
	go fmt ./...
	go vet ./...

test:
	go test -v ./...

migrate-up:
	migrate -path db/migrations -database "mysql://bloger:bloger@tcp(localhost:33306)/bloger" up

migrate-down:
	migrate -path db/migrations -database "mysql://bloger:bloger@tcp(localhost:33306)/bloger" down
