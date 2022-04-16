FROM golang:latest

WORKDIR /app
COPY . /app

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 8080

CMD ["go", "run", "main.go"]