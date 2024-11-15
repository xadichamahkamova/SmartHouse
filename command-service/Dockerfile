FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o command-service ./cmd/main.go

EXPOSE 9090

CMD ["./command-service"]
