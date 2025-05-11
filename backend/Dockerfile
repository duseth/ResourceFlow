# Этап сборки
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Установка необходимых инструментов для сборки
RUN apk add --no-cache git gcc musl-dev

# Копирование и загрузка зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app

# Финальный образ
FROM alpine:3.18

WORKDIR /app

# Копирование бинарного файла из builder
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config
COPY --from=builder /app/migrations ./migrations

# Установка необходимых системных пакетов
RUN apk --no-cache add ca-certificates tzdata

# Создание непривилегированного пользователя
RUN adduser -D -g '' appuser && \
    chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

CMD ["./main"] 