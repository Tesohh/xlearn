include .env

test:
	@go test -v ./...

run:
	@go run .
