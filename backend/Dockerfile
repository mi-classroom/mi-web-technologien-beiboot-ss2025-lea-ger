FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server .

FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache exiftool

COPY --from=builder /app/server .
RUN chmod +x ./server

EXPOSE 8080

CMD ["./server"]
