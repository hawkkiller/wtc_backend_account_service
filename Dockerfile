# syntax=docker/dockerfile:1
# Build Stage
FROM golang:1.17.3-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g main.go --output docs/wtc
RUN go build -o main .



# Run Stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
RUN apk update && apk add bash

COPY wait-for-it.sh .
COPY .dev.env .
RUN chmod +x wait-for-it.sh


EXPOSE 9000
CMD ["./main"]