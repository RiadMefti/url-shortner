# Dockerfile
FROM golang:1.23-alpine

# Install build dependencies for SQLite
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with CGO enabled
ENV CGO_ENABLED=1
RUN go build -o main

# Final stage
FROM alpine:latest

# Install runtime dependencies for SQLite
RUN apk add --no-cache sqlite

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
# Copy static files if you have any
COPY static/ static/

# Create data directory
RUN mkdir -p /data && chmod 777 /data

# Set environment variable for database path
ENV SQLITE_DB_PATH=/data/app.db

EXPOSE 8080

CMD ["./main"]