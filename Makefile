default: all

all: clean test build

.PHONY: clean
clean:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./...

.PHONY: test-cover
test-cover:
	@go test -cover -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: build
build:
	@go build -v ./...