FROM golang:1.24 AS builder
WORKDIR /app
COPY main.go .
RUN GOOS=linux go build -o app main.go

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache libc6-compat
COPY --from=builder /app/app .
RUN chmod +x ./app
EXPOSE 8080 8090 8081
CMD ["./app"]
