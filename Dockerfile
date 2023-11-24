ARG AUTH_API_KEY
ARG AUTH_API_PORT
ARG AUTH_JWT_SECRET
ARG SMTP_HOST
ARG SMTP_PORT
ARG SMTP_LOGIN
ARG SMTP_PASSWORD

FROM golang:1.21.0-alpine  AS builder

ENV AUTH_API_KEY=${AUTH_API_KEY}
ENV AUTH_API_PORT=${AUTH_API_PORT}
ENV AUTH_JWT_SECRET=${AUTH_JWT_SECRET}
ENV SMTP_HOST=${SMTP_HOST}
ENV SMTP_PORT=${SMTP_PORT}
ENV SMTP_LOGIN=${SMTP_LOGIN}
ENV SMTP_PASSWORD=${SMTP_PASSWORD}

COPY . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN set -Eeux && \
    go mod download && \
    go mod verify

EXPOSE 80

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-w -s" \
    -o main

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/main .


ENTRYPOINT ["./main"]

