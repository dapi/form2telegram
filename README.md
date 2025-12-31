# form2telegram

[![Test](https://github.com/dapi/form2telegram/actions/workflows/test.yml/badge.svg)](https://github.com/dapi/form2telegram/actions/workflows/test.yml)
[![Release](https://github.com/dapi/form2telegram/actions/workflows/release.yml/badge.svg)](https://github.com/dapi/form2telegram/actions/workflows/release.yml)

Webhook-мост для пересылки данных из Яндекс.Форм в Telegram.

## Быстрый старт

### Docker

```bash
cp .env.example .env
# Отредактируйте .env файл

docker compose up -d
```

### Локально

```bash
export TELEGRAM_BOT_TOKEN=your_token
export TELEGRAM_CHAT_ID=your_chat_id
go run main.go
```

## API

### POST /yandex-form-webhook

Принимает JSON от Яндекс.Форм:

```json
{
  "answers": [
    {"key": "email", "value": "test@example.com"},
    {"key": "Имя", "value": "Иван"}
  ]
}
```

**Ответы:**

| Код | Описание |
|-----|----------|
| 200 | Сообщение отправлено в Telegram |
| 400 | Невалидный JSON |
| 405 | Метод не POST |
| 500 | Ошибка отправки в Telegram (сеть, неверный токен и т.д.) |

### GET /health

Health-check эндпоинт.

## Настройка Яндекс.Форм

1. Создайте форму в Яндекс.Формах
2. Перейдите в "Интеграции" → "Webhook"
3. Укажите URL: `https://your-domain.com/yandex-form-webhook`

## Переменные окружения

| Переменная | Обязательная | Описание |
|------------|--------------|----------|
| `TELEGRAM_BOT_TOKEN` | Да | Токен Telegram бота |
| `TELEGRAM_CHAT_ID` | Да | ID чата для отправки |
| `PORT` | Нет | Порт сервера (по умолчанию 8080) |

## Логирование

Сервис пишет логи в stdout:

```
2025/01/15 10:30:00 Starting server on :8080
2025/01/15 10:30:05 POST /yandex-form-webhook 45.123ms
2025/01/15 10:30:10 Failed to send message: send request: connection refused
```

**Что логируется:**
- Старт и остановка сервера
- Каждый HTTP-запрос (метод, путь, время выполнения)
- Ошибки парсинга JSON и отправки в Telegram
