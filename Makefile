include .env

test:
	@go test -v ./...

run:
	@go run .

svelte production:
	@cd frontend; PRODUCTION=true bun run index.js