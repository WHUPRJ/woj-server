PROJECT=server

GO := go

LDFLAGS += -X main.BuildTime=$(shell date -u '+%Y-%m-%d-%I-%M-%S')
LDFLAGS += -X main.Version=$(shell cat VERSION)+$(shell git rev-parse HEAD)
LDFLAGS += -s -w

GOBUILD := $(GO) build -o $(PROJECT) -ldflags '$(LDFLAGS)' ./cmd/app

.PHONY: all build clean run dep swagger

default: all

all: clean build

build: swagger dep
	$(GOBUILD)

clean:
	rm -f $(PROJECT)

run: clean swagger dep build
	./$(PROJECT) run

dep:
	go mod tidy && go mod download

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g internal/router/api.go -o internal/router/docs
