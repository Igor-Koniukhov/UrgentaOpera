FROM golang:1.18-alpine as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["go", "run", "./cmd/main.go"]

