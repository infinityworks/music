.PHONY: build test run

build:
	go build -o music cmd/music/main.go

test:
	go test ./...

run:
	go run cmd/music/main.go
