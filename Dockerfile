# Stage 1: Build the Go application
FROM golang:1.22 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o selectstar ./main.go

# Stage 2: Create a minimal image for running the application
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/selectstar .
RUN touch /app/queries.db

# Expose the port that your application will run on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["./selectstar"]
