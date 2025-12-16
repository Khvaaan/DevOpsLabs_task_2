# DevOps Labs — News App

Репозиторий содержит лабораторные работы по курсу DevOps / Docker / Security.

Проект включает:
- настройку и базовое укрепление Linux-хоста (Lab 1);
- разработку REST API приложения на Go (Lab 2);
- контейнеризацию приложения, настройку HTTPS и развёртывание GitLab (Lab 3).

---

## Общая информация

- ОС: Ubuntu 22.04 LTS  
- Архитектура: amd64  
- Язык приложения: Go  
- Контейнеризация: Docker, Docker Compose  
- HTTPS: self-signed сертификаты  
- Хранение секретов: Docker secrets (file-based)

---

## Структура репозитория

```
deploy/
├── app/
│ ├── docker-compose.yml
│ ├── entrypoint.sh
│ ├── schema.sql
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


Каталоги `secrets`, а также runtime-данные GitLab (`data`, `config`, `logs`) не добавляются в git.

---

## Lab 1 — Host Linux / Security

В рамках лабораторной выполнено:
- обновление системы;
- настройка SSH (нестандартный порт);
- настройка firewall (UFW);
- ограничение прав доступа к системным файлам;
- настройка лимитов пользователей;
- настройка HTTPS с self-signed сертификатом и проверка через `curl`.

---

## Lab 2 — REST API (Go)

Реализовано REST API с CRUD-операциями.

Поддерживаемые хранилища:
- in-memory (memdb);
- PostgreSQL;
- MongoDB (подготовлено).

Проверка локального запуска:
```
go test ./...
go run ./cmd/server
```

## ab 3 — Docker / HTTPS / GitLab
### Приложение (nginx + app + PostgreSQL)

Приложение разворачивается с использованием docker-compose и доступно по HTTPS.

При первом старте PostgreSQL схема БД автоматически инициализируется
через schema.sql, подключённый в /docker-entrypoint-initdb.d/.

Запуск стека приложения:
```
cd deploy/app
docker compose up -d
```

Проверка записи данных в БД через API:
```
curl -k -X POST https://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"CHECK","done":false}'
```

Ожидаемый результат — HTTP 200 и появление записи в таблице tasks.

### GitLab под HTTPS

GitLab разворачивается отдельным docker-compose файлом и работает по HTTPS
с использованием self-signed сертификата.

Запуск GitLab:
```
cd deploy/gitlab
docker compose up -d
```

Проверка доступности:
```
curl -k -I https://localhost
curl -k -I https://localhost/users/sign_in

```

GitLab возвращает HTTP 302 с редиректом на /users/sign_in, что является
штатным поведением для неаутентифицированного пользователя и подтверждает
корректную работу HTTPS.
Соответственно, при обращении на "localhost/users/sign_in" отдаёт 200.

### Работа с секретами

Пароли БД и TLS-ключи не хранятся в открытом виде в compose-файлах.

Все чувствительные данные передаются через docker secrets.

Каталоги ```deploy/*/secrets``` исключены из git.