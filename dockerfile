FROM golang:1.20.2-alpine3.17 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -o main.app .

FROM alpine:latest

COPY --from=builder /app/main.app .

CMD ["./main.app"]