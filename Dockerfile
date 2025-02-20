FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

# Use a minimal Debian image for the final container
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

COPY migrations ./migrations

RUN chmod +x /app/main

EXPOSE 8080

CMD ["/app/main"]
