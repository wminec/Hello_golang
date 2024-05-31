FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

RUN adduser -u 185 -g 0 -d /home/go -s /sbin/nologin go && \
    mkdir -p /app/main && \
    chown -R go:root /app/main && \
    chmod -R 775 /app

WORKDIR /app/main

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
