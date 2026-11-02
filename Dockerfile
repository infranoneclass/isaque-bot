FROM golang:1.27.3 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/isaque-bot ./cmd/bot

FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && adduser -D -H -u 10001 appuser

WORKDIR /app
COPY --from=builder /out/isaque-bot ./isaque-bot

USER appuser
ENTRYPOINT ["./isaque-bot"]
