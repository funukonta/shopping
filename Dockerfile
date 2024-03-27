# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine

WORKDIR /shopping

COPY . .

# Download dependencies
RUN go mod download

#build
RUN go build -v -o /shopping/shoppingApp ./cmd/main.go

# Expose port 8080
EXPOSE 8080

#run app
ENTRYPOINT [ "/shopping/shoppingApp" ]