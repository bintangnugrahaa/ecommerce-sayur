FROM golang:1.25.5-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main main.go

EXPOSE 8082

CMD ["./main", "start"]