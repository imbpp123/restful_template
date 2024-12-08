FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/api

#------------------------------
FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 3333

CMD ["./app"]