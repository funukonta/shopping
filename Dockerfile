# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine

WORKDIR /shopping

COPY . .

# Download dependencies
RUN go mod download

#build
RUN go build -v -o /shopping/shoppingApp ./cmd/main.go

#run app
ENTRYPOINT [ "/shopping/shoppingApp" ]