FROM golang:1.24.3-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bot ./cmd

FROM alpine:latest
COPY --from=builder /app/bot /bot
CMD ["/bot"]