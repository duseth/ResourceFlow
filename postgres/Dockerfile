FROM timescale/timescaledb:latest-pg14

# Копируем инициализационные скрипты
COPY ./backend/migrations/*.sql /docker-entrypoint-initdb.d/

# Устанавливаем права на выполнение скриптов
RUN chmod +x /docker-entrypoint-initdb.d/*.sql 