FROM golang:1.17.6

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o servers_distributeApp ./main.go

RUN chmod +x /app/servers_distributeApp

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/servers_distributeApp /app
CMD ["/app/servers_distributeApp"]
