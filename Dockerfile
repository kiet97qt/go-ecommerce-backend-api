FROM golang:1.25.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server .
COPY --from=builder /app/configs ./configs

ENV APP_ENV=dev
ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["./server"]


