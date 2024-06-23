FROM golang:1.22 as builder

WORKDIR /app

#COPY go.mod go.sum ./
COPY . .
RUN go mod download



RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./client

FROM alpine:latest

RUN apk --no-cache add -U tzdata
RUN adduser -u 185 -G root -h /home/go -s /sbin/nologin -D go && \
    mkdir -p /app/main

WORKDIR /app/main

COPY --from=builder /app/main .

RUN chown -R go:root /app && \
    chmod -R 755 /app/main

USER 185

ENTRYPOINT ["./main"]
