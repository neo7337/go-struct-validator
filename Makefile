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

.PHONY: test-bench
test-bench:
	@go test -run=Bench -bench=. -benchtime 5000000x

.PHONY: build
build:
	@go build -v ./...