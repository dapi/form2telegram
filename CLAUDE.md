# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Описание проекта

Webhook-мост для приёма данных от Яндекс.Форм и пересылки в Telegram-чат.

## Команды разработки

```bash
# Запуск
go run main.go

# Сборка
go build -o form2telegram .

# Тесты
go test ./...

# Тесты с покрытием
go test -cover ./...

# Один тест
go test -run TestHandleWebhook ./internal/handler/

# Линтер
golangci-lint run

# Docker
docker build -t form2telegram .
docker run -p 8080:8080 --env-file .env form2telegram
```

## Архитектура

```
main.go                 # Точка входа: инициализация сервера и graceful shutdown
internal/
├── handler/           # HTTP-обработчики (POST /yandex-form-webhook, GET /health)
├── telegram/          # Клиент Telegram Bot API
└── formatter/         # Форматирование формы в Markdown-сообщение
```

### Поток данных

1. Яндекс.Формы отправляет POST на `/yandex-form-webhook`
2. `handler` парсит JSON и передаёт в `formatter`
3. `formatter` преобразует данные формы в Markdown
4. `telegram` клиент отправляет сообщение в чат

## Конфигурация

Переменные окружения:
- `TELEGRAM_BOT_TOKEN` — токен бота (обязательно)
- `TELEGRAM_CHAT_ID` — ID чата для отправки (обязательно)
- `PORT` — порт сервера (по умолчанию 8080)

## Технологии

- Go 1.24+ (mise)
- Стандартная библиотека `net/http`
- Telegram Bot API (без сторонних SDK)
