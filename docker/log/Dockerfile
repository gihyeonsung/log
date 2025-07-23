# builder
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY . .

RUN apk add --no-cache \
    gcc \
    musl-dev \
    git \
    make \
    bash && \
    go build -o log .

# runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/log ./log

CMD ["./log"]
