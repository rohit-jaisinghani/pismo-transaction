#build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY vendor/ ./vendor

COPY . .

RUN go build -mod=vendor -o app .

#runtime stage
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
