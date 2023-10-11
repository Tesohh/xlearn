include .env

test:
	@go test -v ./...

run:
	@go run .

svelte:
	@cd frontend; npm run dev