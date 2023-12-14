FROM golang:1.21 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/grpc_server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
RUN chmod +x ./app
ENTRYPOINT ["./app"]