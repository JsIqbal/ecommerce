# Build stage
FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

# Run final stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080

# Command to run the executable
CMD ["/app/main", "serve-rest"]
