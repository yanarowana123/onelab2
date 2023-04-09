FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o cmd/main .

EXPOSE 8081

CMD ["./main"]
