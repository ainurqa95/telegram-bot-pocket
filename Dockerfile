FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY . /github.com/ainurqa95/telegram-bot-pocket/
WORKDIR /github.com/ainurqa95/telegram-bot-pocket/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/ainurqa95/telegram-bot-pocket/.bin/bot .
COPY --from=0 /github.com/ainurqa95/telegram-bot-pocket/configs configs/

EXPOSE 80

CMD ["./bot"]