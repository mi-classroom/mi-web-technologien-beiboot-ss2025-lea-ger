FROM golang:1.24-alpine

ENV GO111MODULE=on

WORKDIR /app

RUN apk add --no-cache git
RUN apk add --no-cache exiftool

RUN which exiftool

RUN exiftool -ver

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV PATH="/usr/bin:$PATH"

RUN migrate -path ./migrations -database "postgres://admin:admin@localhost:5432/iptc-editor?sslmode=disable" up

CMD ["go", "test", "./..."]