# DevOps Labs — News App

Репозиторий содержит лабораторные работы по курсу DevOps / Docker / Security.

Проект включает:
- настройку и базовое укрепление Linux-хоста (Lab 1),
- разработку REST API приложения на Go (Lab 2),
- контейнеризацию приложения, настройку HTTPS и развёртывание GitLab (Lab 3).

---

## Общая информация

- ОС: Ubuntu 22.04 LTS  
- Архитектура: amd64  
- Язык приложения: Go  
- Контейнеризация: Docker, Docker Compose  
- HTTPS: self-signed сертификаты  
- Работа с секретами: Docker secrets (file-based)

---

## Структура репозитория
```
deploy/
├── app/
│ ├── docker-compose.yml
│ ├── entrypoint.sh
│ ├── nginx/
│ │ └── nginx.conf
│ └── secrets/
│ ├── pg_password.txt
│ ├── tls.crt
│ └── tls.key
└── gitlab/
├── docker-compose.yml
├── config/
├── data/
├── logs/
└── secrets/
├── root_password.txt
├── tls.crt
└── tls.key
```


Каталоги `secrets` содержат чувствительные данные и не должны добавляться в git.

---

## Lab 1 — Host Linux / Security

Выполнено:
- обновление системы;
- настройка SSH (нестандартный порт);
- включение и настройка UFW;
- ограничение прав доступа к системным файлам;
- настройка лимитов пользователей;
- настройка HTTPS с self-signed сертификатом и проверка через `curl`.

---

## Lab 2 — REST API (Go)

> Файлы 2й лабы находятся в коммите "upd_1 README"

Реализовано REST API с CRUD-операциями.

Поддерживаемые хранилища:
- in-memory (memdb);
- PostgreSQL;
- MongoDB.

Проверка локального запуска:
```bash
go test ./...
go run ./cmd/
```

## Lab 3 — Docker / HTTPS / GitLab
### Приложение (nginx + app + PostgreSQL)

Запуск стека приложения:
```
cd deploy/app
docker compose up -d
```

Проверка записи в БД по HTTPS:
```
curl -k -X POST https://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"TEST","done":false}'
```
Результат: HTTP 200 и появление записи в PostgreSQL.

## GitLab под HTTPS
GitLab разворачивается отдельным docker-compose
Запуск:
```
cd deploy/gitlab
docker compose up -d
```

Проверка доступности:
```
curl -k -I https://localhost
```
GitLab возвращает HTTP 302 с редиректом на "/users/sign_in", что является корректным поведением для неаутентифицированного пользователя и подтверждает работоспособность HTTPS.
```
curl -k -I https://localhost/users/sign_in
```
При обращении к "https://localhost/users/sign_in" отдаёт 200

### Работа с секретами

Пароли БД и TLS-ключи не хранятся в открытом виде в compose-файлах.

Все секреты подключаются через docker secrets.

Каталоги deploy/*/secrets исключены из git.
