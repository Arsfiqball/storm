FROM golang:1.19-alpine

WORKDIR /app

ENV CGO_ENABLED=0
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

CMD ["air", "-c", ".air.toml"]
