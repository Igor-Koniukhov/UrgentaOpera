FROM golang:1.18-alpine as build

WORKDIR /app/migrations

COPY . /app/migrations/


RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:3.14.6

WORKDIR /app

COPY --from=build /go/bin/goose/ app/migrations



