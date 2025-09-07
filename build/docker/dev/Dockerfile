FROM golang:1.24-alpine

WORKDIR /app

ENV CGO_ENABLED=0
RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest

# Add docker build cache for Go modules
ENV GOMODCACHE=/go/pkg/mod
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p /go/pkg/mod /root/.cache/go-build

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

CMD ["air", "-c", ".air.toml"]
