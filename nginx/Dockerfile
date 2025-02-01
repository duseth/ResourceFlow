FROM nginx:alpine

# Копирование конфигурации
COPY conf.d/default.conf /etc/nginx/conf.d/default.conf

# Создание директории для SSL сертификатов
RUN mkdir -p /etc/nginx/ssl

# Настройка прав доступа
RUN chown -R nginx:nginx /etc/nginx/ssl && \
    chmod -R 600 /etc/nginx/ssl

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"] 