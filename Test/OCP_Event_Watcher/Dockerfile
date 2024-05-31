FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

RUN adduser -u 185 -g 0 -d /home/go -s /sbin/nologin go && \
    mkdir -p /app/main

WORKDIR /app/main

COPY --from=builder /app/main .

RUN chown -R go:root /app && \
    chmod -R 755 /app/main

USER 185

ENTRYPOINT ["./main"]
