FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go

EXPOSE 8090

CMD ["./main", "start"]