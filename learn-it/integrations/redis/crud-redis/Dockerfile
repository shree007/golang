FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o crud-redis
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /crud-redis/
COPY --from=builder /app/crud-redis .
CMD ["./crud-redis"]
