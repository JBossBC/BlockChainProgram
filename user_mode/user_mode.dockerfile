FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o userApp ./main.go

RUN chmod +x /app/userApp

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/userApp /app
CMD ["/app/userApp"]
