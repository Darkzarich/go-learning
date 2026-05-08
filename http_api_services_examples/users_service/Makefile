.PHONY: build run-api run-cron

build:
	GOOS=linux go build -o bin/api cmd/api/main.go
	GOOS=linux go build -o bin/cronjobs cmd/cronjobs/main.go

run-api:
	go run cmd/api/main.go

run-cron:
	go run cmd/cronjobs/main.go