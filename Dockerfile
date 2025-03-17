# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder
RUN apt-get update && apt-get install -y gcc libsqlite3-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o block1_back main.go

FROM debian:latest
RUN apt-get update && apt-get install -y sqlite3
WORKDIR /app
COPY --from=builder /app/block1_back .
COPY data.db ./data.db
EXPOSE 9090
CMD ["./block1_back"]
