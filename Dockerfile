## Build
FROM golang:1.21.0-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod vendor

RUN env GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -o app

## Deploy
FROM alpine

WORKDIR /

COPY --from=builder /build .

EXPOSE 8080

RUN chmod +x ./app

CMD ["./app"]