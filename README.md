# ResourceFlow

ResourceFlow - это современное решение для мониторинга и оптимизации серверных ресурсов, разработанное в рамках магистерской диссертации. Система предоставляет комплексный подход к мониторингу серверных ресурсов с интуитивно понятным веб-интерфейсом и расширенной аналитикой в реальном времени.

## Ключевые особенности

- Real-time мониторинг системных ресурсов
- Интерактивные дашборды с визуализацией метрик
- Автоматическая оптимизация ресурсов
- Исторический анализ и прогнозирование
- Система оповещений и алертов
- Мониторинг нескольких серверов

## Технологии

- Backend: Go 1.21+
- Frontend: React + TypeScript
- База данных: PostgreSQL + TimescaleDB
- Кэширование: Redis
- Контейнеризация: Docker + Docker Compose
- CI/CD: GitHub Actions

## Быстрый старт

1. Клонируйте репозиторий:
```bash
git clone https://github.com/duseth/ResourceFlow.git
cd ResourceFlow
```

2. Скопируйте файл с переменными окружения:
```bash
cp .env.example .env
```

3. Запустите проект:
```bash
make dev
```

4. Откройте в браузере:
- Веб-интерфейс: http://localhost:80
- Swagger документация: http://localhost:80/api/docs

## Документация

Полная документация доступна в директории `docs/`:
- [Начало работы](docs/user/getting-started.md)
- [API документация](docs/api/swagger/swagger.json)
- [Архитектура](docs/architecture/overview.md)
- [Руководство разработчика](docs/development/setup.md)

## Разработка

```bash
# Сборка проекта
make build

# Запуск тестов
make test-backend
make test-frontend

# Проверка кода
make lint-backend
make lint-frontend

# Генерация документации
make docs
```

## Лицензия

MIT License - см. [LICENSE](LICENSE) файл для деталей.
