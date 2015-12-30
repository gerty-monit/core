.PHONY: install build test run clean-ts

build:
	@go build -v ./...
	@go vet -v ./...

install:
	go get github.com/tools/godep
	godep restore

test:
	@go test ./...

clean-ts:
	rm -rf public/js/compiled

ts_files = $(shell find public/typescript -name '*.tsx')
ts-compile: $(ts_files)
	@tsc --noImplicitAny --jsx react --rootDir public/typescript --outDir public/js/compiled $?

ts-compile-watch: $(ts_files)
	@tsc -w --noImplicitAny --jsx react --rootDir public/typescript --outDir public/js/compiled $?

run: ts-compile
	@go run *.go

test-cover:
	@echo "mode: set" > acc.coverage-out
	@go test -coverprofile=services.coverage-out ./services
	@cat services.coverage-out | grep -v "mode: set" >> acc.coverage-out
	@go tool cover -html=acc.coverage-out
	@rm *.coverage-out

run-prod: ts-compile
	@GIN_MODE=release go run main.go
