FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache exiv2

COPY --from=builder /app/server .
RUN chmod +x ./server

EXPOSE 8080

CMD ["./server"]
