FROM golang:1.19-alpine

WORKDIR /app

ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

# Build binary to use in command later
RUN go build -o ./bin/server ./cmd/server/main.go

# Run the server binary
CMD ["./bin/server"]
