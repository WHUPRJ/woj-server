GO := go

LDFLAGS += -X cmd.BuildTime=$(shell date -u '+%Y-%m-%d-%I-%M-%S')
LDFLAGS += -X cmd.Version=$(shell cat VERSION)+$(shell git rev-parse HEAD)
LDFLAGS += -s -w

GOBUILD := $(GO) build -ldflags '$(LDFLAGS)'
GOBIN   := $(shell go env GOPATH)/bin

.PHONY: all server runner build clean dep swagger fmt

default: all

all: clean build

server: swagger dep
	$(GOBUILD) -o server ./cmd/server

runner: swagger dep
	$(GOBUILD) -o runner ./cmd/runner

build: runner server

clean:
	rm -f runner
	rm -f server

dep:
	go mod tidy && go mod download

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	$(GOBIN)/swag init -g internal/router/api.go -o internal/router/docs

fmt:
	go fmt ./...
