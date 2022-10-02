default: all

all: clean test build

.PHONY: clean
clean:
	@go mod tidy

.PHONY: test
test:
	go test -cover -coverprofile=coverage.txt -covermode=atomic ./...
