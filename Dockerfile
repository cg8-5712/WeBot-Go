FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o WeBot bot.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/WeBot .
COPY .env .
ENTRYPOINT ["./WeBot"]
