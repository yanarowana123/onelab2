FROM golang:1.20 AS stage1
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd

FROM alpine:latest as stage2
WORKDIR /app
COPY --from=stage1 /app/app .
CMD ["./app"]