.PHONY: build

build:
	export GO111MODULE=on
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/api-ping-get pkg/api/ping/get/main.go
