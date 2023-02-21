## Установка пакетов

    $ go mod download

## Настройка конфигурации
Нужно создать файл `.env` и добавить переменные:

    # DATABASE
    PSQL_USER="postgres"
    PSQL_PASSWORD="postgres"
    PSQL_PORT="5432"
    PSQL_DATABASE="xenforo"
    
    # PG_ADMIN
    PG_ADMIN_EMAIL="admin@admin.ru"
    PG_ADMIN_PASSWORD="admin"