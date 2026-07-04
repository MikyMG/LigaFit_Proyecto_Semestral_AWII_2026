FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/ligafit-api ./cmd/api

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -u 10001 appuser

WORKDIR /app

COPY --from=builder /bin/ligafit-api /app/ligafit-api

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/ligafit-api"]