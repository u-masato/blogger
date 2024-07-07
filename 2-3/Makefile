# Makefile

.PHONY: run lint test

run:
	go run ./cmd/app/main.go

lint:
	go fmt ./...
	go vet ./...

test:
	go test -v ./...
