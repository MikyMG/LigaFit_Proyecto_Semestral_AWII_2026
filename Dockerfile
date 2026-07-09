# ---------- Etapa de compilación ----------
FROM golang:1.26.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ligafit ./cmd/api

# ---------- Imagen final ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ligafit .

EXPOSE 8080

CMD ["./ligafit"]