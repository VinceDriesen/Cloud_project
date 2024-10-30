FROM bitnami/laravel:11-debian-12

WORKDIR /app

COPY ./app/hospitalWebsite /app

EXPOSE 8000

CMD ["php", "artisan", "serve", "--host=0.0.0.0", "--port=8000"]
