NAME := go_tutorial
VERSION := v0.1.0

SRCS := $(shell find . -type f -name "*.go")

bin/hello_world: $(SRCS)
	go build -o bin/hello_world cmd/hello_world.go

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: test
test:
	go test ./...
