## Build
FROM golang:1.19.0-alpine AS builder

WORKDIR /build

COPY . .

RUN GOOS=linux go build -mod=vendor -ldflags="-s -w" -o app

## Deploy
FROM alpine

WORKDIR /

COPY --from=builder /build .

EXPOSE 8080

USER nonroot:nonroot

CMD ["./app"]