# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0

# Add docker build cache for Go modules
ENV GOMODCACHE=/go/pkg/mod
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p /go/pkg/mod /root/.cache/go-build

# Copy only the dependency files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the application
RUN go build -o ./bin/server ./cmd/server

# Production stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates in case the application makes HTTPS calls
RUN apk --no-cache add ca-certificates

# Copy only the compiled binary from the build stage
COPY --from=builder /app/bin/server ./bin/server

# Ensure the binary is executable
RUN chmod +x ./bin/server

# Create a non-root user to run the application
RUN adduser -D -H -h /app appuser
USER appuser

# Run the server binary
CMD ["./bin/server"]
