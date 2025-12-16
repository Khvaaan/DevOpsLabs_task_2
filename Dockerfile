# syntax=docker/dockerfile:1

# --- build stage: собираем бинарник ---
FROM golang:1.22-alpine AS builder
# Рабочая директория внутри контейнера сборки
WORKDIR /src

# Ставим корневые сертификаты и tzdata (иногда нужно для HTTPS/времени)
RUN apk add --no-cache ca-certificates tzdata

# Копируем файлы модулей отдельно, чтобы кешировать зависимости
COPY go.mod go.sum ./
# Скачиваем зависимости (кешируется слоем)
RUN go mod download

# Копируем весь исходный код проекта
COPY . .

# Отключаем CGO для статического бинарника (проще для минимального runtime)
ENV CGO_ENABLED=0
# Целевая ОС сборки — Linux
ENV GOOS=linux
# Архитектура (Ubuntu VM обычно amd64)
ENV GOARCH=amd64

# Собираем бинарник из ./cmd/server
RUN go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/server


# --- runtime stage: минимальный образ для запуска ---
FROM alpine:3.20
# Рабочая директория приложения
WORKDIR /app

# Сертификаты/таймзона для HTTPS-запросов и корректного времени
RUN apk add --no-cache ca-certificates tzdata

# Создаём непривилегированного пользователя (запуск не от root)
RUN adduser -D -H -s /sbin/nologin appuser

# Копируем собранный бинарник из build stage
COPY --from=builder /out/app /app/app

# Отдаём права на каталог пользователю приложения
RUN chown -R appuser:appuser /app

# Запускаем приложение не от root
USER appuser

# Документируем порт, который слушает приложение
EXPOSE 8080

# Команда запуска контейнера
ENTRYPOINT ["/app/app"]
