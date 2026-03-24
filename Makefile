APP_NAME=GolangStore

.PHONY: help
help:
	@go version

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test ./...
