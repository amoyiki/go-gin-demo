.PHONY: build
all: build
build:
	CGO_ENABLED=0 go build -o go-gin-demo cmd/main.go
