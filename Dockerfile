# ------------------------------
# Build stage
# ------------------------------
FROM golang:1.25-alpine AS builder

# Install build tools
RUN apk add --no-cache git make bash

# Set working dir
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the server
RUN go build -o server ./cmd/server

# ------------------------------
# Runtime stage
# ------------------------------
FROM alpine:3.20

# Install Postgres client (optional, useful for debugging/migrations)
RUN apk add --no-cache ca-certificates postgresql-client

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .

# Expose ports (HTTP :8080, gRPC :9090)
EXPOSE 8080 9090

# Set environment variable for DB connection (override in compose)
ENV DATABASE_URL=postgres://postgres:postgres@db:5432/notes?sslmode=disable

CMD ["./server"]
