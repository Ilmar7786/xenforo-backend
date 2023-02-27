## Установка пакетов

    $ go mod download

## Настройка конфигурации
Нужно создать файл `.env` и добавить переменные:

    # DATABASE
    PSQL_USER="postgres"
    PSQL_PASSWORD="postgres"
    PSQL_HOST="localhost"
    PSQL_PORT="5432"
    PSQL_DATABASE="xenforo"
    
    # Mail
    MAIL_FROM="fromEmail@yandex.ru"
    MAIL_PASSWORD="password"
    MAIL_USERNAME="username"
    MAIL_HOST="smtp.yandex.ru"
    MAIL_PORT=465
    #MAIL_SSL=bool (default: true)

    # FlashLiveSports API
    FLASH_LIVE_SPORTS_TOKEN="token"

    # (FOR DOCKER)
    # PG_ADMIN
    PG_ADMIN_EMAIL="admin@admin.ru"
    PG_ADMIN_PASSWORD="admin"