FROM golang:latest AS builder

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redis-proxy redis-proxy.go

FROM debian:buster-slim

COPY --from=builder /app/redis-proxy /

RUN groupadd -g 999 appuser && \
    useradd -r -u 999 -g appuser appuser

EXPOSE 6379

USER 999:999

CMD ["/redis-proxy"]
