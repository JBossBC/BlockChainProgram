FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY verify/. /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./main.go

RUN chmod +x /app/brokerApp

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/brokerApp /app
CMD ["/app/brokerApp"]
