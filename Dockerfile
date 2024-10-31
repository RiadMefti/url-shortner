# Use the official Golang image as the base image
FROM golang:1.23.2-alpine AS builder


# Install build dependencies for SQLite and CGO
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with CGO enabled for SQLite support
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Use a smaller alpine image for the final container
FROM alpine:latest

RUN apk add --no-cache sqlite-libs

# Copy everything from builder
WORKDIR /app
COPY --from=builder /app .
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]