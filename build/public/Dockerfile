# Build stage
FROM golang:1.19-alpine3.16 AS builder
RUN apk add --no-cache git gcc musl-dev
WORKDIR /app
COPY .. .
RUN go mod vendor
RUN go build -o main -tags musl cmd/public/main.go

# Run stage
FROM alpine:3.16 as runner
WORKDIR /app
COPY ../../static static
COPY --from=builder /app/main .