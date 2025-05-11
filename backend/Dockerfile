FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN migrate -path ./migrations -database "postgres://admin:admin@localhost:5432/iptc-editor?sslmode=disable" up

RUN GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache exiftool

COPY --from=builder /app/server .
RUN chmod +x ./server

EXPOSE 8080

CMD ["./server", "-bind", "0.0.0.0:8080"]
