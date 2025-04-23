FROM golang:1.23.5-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY templates ./templates

RUN go install github.com/google/wire/cmd/wire@latest
RUN go generate ./...

RUN go build -o main ./cmd

FROM alpine:3.21.3

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser

COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates

RUN chown -R appuser:appgroup /home/appuser

USER appuser

EXPOSE 8080

CMD ["./main"]
