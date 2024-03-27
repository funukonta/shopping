# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine AS builder

WORKDIR /shopping

COPY . .

# Download dependencies
RUN go mod download

# Build the Go binary
RUN go build -v -o /shopping/shoppingApp ./cmd/main.go

# Stage 2: Create a minimal image to run the binary
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /shopping/shoppingApp .

# Expose port 8080
EXPOSE 8080

# Command to run the binary
CMD ["./shoppingApp"]
