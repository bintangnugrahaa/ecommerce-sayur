FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go
RUN go build -o worker-order cmd/worker-order.go
RUN go build -o worker-update-payment-order cmd/worker-update-payment-order.go

EXPOSE 8083

CMD ["./main", "start"]