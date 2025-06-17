# 🏦 Bank REST API (Go)

## 📋 Описание
Банковский сервис, реализованный на Go. Поддерживает регистрацию, авторизацию (JWT), создание счетов, выпуск карт, переводы, кредиты, аналитику и интеграцию с ЦБ РФ и SMTP.

## 🚀 Возможности
- Регистрация и вход
- JWT-аутентификация
- Управление банковскими счетами
- Генерация карт (Лун, PGP, HMAC)
- Переводы между счетами
- Кредитование и графики платежей
- Автоматическое списание задолженностей
- Финансовая аналитика
- Интеграция с ЦБ РФ (SOAP)
- Email-уведомления (SMTP)

## 🧱 Стек
- Язык: Go 1.23+
- БД: PostgreSQL 17
- SMTP: любой TLS-поддерживаемый сервер
- SOAP: ЦБ РФ — https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx

## 📦 Сборка и запуск

```bash
docker-compose up --build
```

## 🛠 Makefile

```bash
make build     # Собрать бинарник
make run       # Запустить локально
make test      # Прогнать тесты
```

## 📬 Переменные окружения (.env)

```
JWT_SECRET=your_jwt_secret
SMTP_HOST=smtp.example.com
SMTP_USER=noreply@example.com
SMTP_PASS=super_secret
PGP_PUBLIC_KEY=...
PGP_PRIVATE_KEY=...
```

## 📚 Автор
Учебный проект: реализация backend для банковского сервиса.
