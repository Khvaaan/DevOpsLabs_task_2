#!/bin/sh
set -e

# Читаем пароль из docker secret
PG_PASS="$(cat /run/secrets/pg_password)"

# Собираем строку подключения
export DATABASE_URL="postgres://newsuser:${PG_PASS}@postgres:5432/news?sslmode=disable"

# Запускаем приложение
exec /app/app
