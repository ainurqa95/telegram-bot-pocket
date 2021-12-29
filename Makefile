.PHONY:
.SILENT:

build: 
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

docker-image:
	docker build -t telegram-bot-pocket:v1.0 .

start-container:
	docker run --env-file .env -p 80:80 telegram-bot-pocket:v1.0
