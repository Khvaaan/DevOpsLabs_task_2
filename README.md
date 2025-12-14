# DevOpsLabs — Lab 2: Task CRUD (Go)

REST API сервис с 4 методами (GET/POST/PUT/DELETE) и 3 реализациями хранилища:
- **memdb** (fake/in-memory) — используется для проверки
- **postgres** — реализация интерфейса на будущее
- **mongo** — реализация интерфейса на будущее

## Requirements
- Go **1.18+**

## Run (Fake DB)
Из корня проекта:
```
go mod tidy
go run ./cmd/server
```

Сервис стартует на:

http://127.0.0.1:8080

## API

Endpoints:

GET /tasks — получить список задач

POST /tasks — создать задачу

PUT /tasks — обновить задачу

DELETE /tasks — удалить задачу

## Examples (curl)

Get:
curl http://127.0.0.1:8080/tasks

Create:
curl -X POST http://127.0.0.1:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":10,"title":"t","done":false,"created_at":0}'

Update:
curl -X PUT http://127.0.0.1:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":10,"title":"t2","done":true,"created_at":0}'

Delete:
curl -X DELETE http://127.0.0.1:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":10}'

## Проверка, что все пакеты собираются:

go test ./...

## Файл schema.sql содержит схему таблицы tasks (на будущее для PostgreSQL).
