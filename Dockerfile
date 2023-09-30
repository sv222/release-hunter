# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 go build -o rh -a -installsuffix cgo ./cmd/release-hunter

# Stage 2: Create a minimal image with the binary
FROM alpine:latest

# Set the working directory in the new stage
WORKDIR /app

# Copy the built binary from the builder stage into this stage
COPY --from=builder /app/rh .

# Make the binary executable
RUN chmod +x rh

# Set environment variables (if needed)
# ENV GITHUB_TOKEN=

# Define the command to run your application
ENTRYPOINT ["./rh"]