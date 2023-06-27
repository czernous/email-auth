FROM golang:1.20.5-alpine  AS builder

COPY . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN set -Eeux && \
    go mod download && \
    go mod verify


RUN GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-w -s" \
    -o main

CMD [ "./main" ]

# FROM scratch

# WORKDIR /root/

# COPY --from=builder /app/main .


# ENTRYPOINT ["./main"]

