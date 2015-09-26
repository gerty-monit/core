.PHONY: install build test run

build:
	@go build -v ./...

install:
	@godep restore

test:
	@go test ./...