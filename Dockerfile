FROM golang:1.23.5-alpine AS builder

COPY . /var/www
WORKDIR /var/www

RUN go build -o bin cmd/main.go

FROM alpine:latest
ARG GRPC_PORT

COPY --from=builder . .
WORKDIR /var/www

EXPOSE $GRPC_PORT

CMD ["./bin/main"]