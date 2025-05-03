# user-service

Сервис для обработки запросов в бд на Go с использованием gRPC.

## Требования

- Go 1.24.2 или выше
- PostgreSQL 15 или выше
- Make (опционально)

## Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/LimeOnTop/user-service.git
cd user-service
```

2. Установите зависимости:
```bash
go mod download
```

3. Измените переменные окружения в config/config.yaml

## Запуск приложения

### Подготовка базы данных

1. Примените миграции:
```bash
make migrate-up
```

### Если вы используете Docker:
```bash
docker-compose up --build
```

### Запуск сервера

1. В режиме разработки:
```bash
make run-dev
```

2. В production режиме:
```bash
make run
```

## Доступные команды

- `make run` - запуск приложения
- `make run-dev` - запуск в режиме разработки
- `make migrate-up` - применение миграций
- `make migrate-down` - откат миграций
- `make migrations-create name=<name>` - создание новой миграции
- `make proto-generate` - генерация gRPC контрактов

## Тестирование

Для запуска тестов:
```bash
go test ./...
```

Для проверки покрытия тестами:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## CI/CD

Проект использует GitLab CI/CD для автоматизации сборки и тестирования. Пайплайн включает:
- Сборку приложения
- Запуск тестов
- Проверку покрытия кода тестами (минимум 30%)

## Структура проекта

```
├── cmd/                    # Точки входа приложения
│   ├── app/               # Основное приложение
│   ├── migrate/           # Миграции базы данных
├── config/                # Конфигурация
├── gen/                   # Сгенерированные файлы
├── internal/              # Внутренние пакеты
│   ├── app/               # Точка входа (откуда идёт запуск)
│   ├── adapter/           # Адаптеры к инфре
│   ├── controller/        # Контроллеры (transport)
│   ├── repository/        # Репозитории
│   └── usecase/           # Сценарии использования
├── migrations/            # SQL миграции
└── proto/                 # gRPC протофайлы
```

## API

### gRPC методы

- `Register` - регистрация нового пользователя
- `Login` - авторизация пользователя
- `RefreshToken` - обновление токена
- `ValidateToken` - проверка токена
- `Logout` - выход из системы