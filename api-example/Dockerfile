# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o exe

FROM alpine:3.18.4
WORKDIR /app
COPY --from=builder /app/exe /app
ENTRYPOINT ["/app/exe"]

