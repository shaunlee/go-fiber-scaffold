MAKEFLAGS += -s

.PHONY: dev build run

dev:
	go run -v ./cmd/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/main -ldflags="-s -w" cmd/main.go \
			 && upx bin/main

run:
	docker run --rm -it -w /www -v ${PWD}:/www --net=host debian:12-slim /www/bin/main
