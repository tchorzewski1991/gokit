lint:
	@golangci-lint run -v -c golangci.yaml

test:
	@go test ./...

tidy:
	@go mod tidy
	@go mod verify

PHONY: lint test tidy
