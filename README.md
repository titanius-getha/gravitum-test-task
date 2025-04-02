## Секреты

Создайте на уровень выше репозитория файл `secrets.env` с содержимым:

```env
DB_PASSWORD=<ваш пароль базы данных>
```

## Тестирование

```bash
# Все тесты
go test -v ./...

# Только unit-тесты
go test -v -short ./...

# Интеграционные тесты (требуют Docker)
go test -v -run Integration ./...
```

## Запуск

```bash
# Режим разработки
./scripts/dev.up.sh

# Прод режим
./scripts/prod.up.sh
```
