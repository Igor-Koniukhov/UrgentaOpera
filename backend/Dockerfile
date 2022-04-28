FROM golang:1.18-alpine as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
EXPOSE 8000

CMD ["go", "run", "./cmd/main.go"]

