Event Driven Order System

Сервис обработки заказов, построенный на событийно-ориентированной архитектуре (Event Driven Architecture).

Проект демонстрирует работу backend-сервиса на Go, использующего Kafka для передачи событий и PostgreSQL для хранения данных.

⸻

Архитектура системы

Система работает по следующему принципу:

Клиент
   ↓
HTTP API
   ↓
Kafka Producer
   ↓
Kafka Topic (orders)
   ↓
Worker Consumer
   ↓
PostgreSQL

Поток обработки заказа
 1. Клиент отправляет HTTP запрос на создание заказа
 2. API принимает запрос
 3. API отправляет событие в Kafka
 4. Kafka сохраняет событие в топик orders
 5. Worker читает сообщение из Kafka
 6. Worker обрабатывает событие
 7. Данные сохраняются в PostgreSQL

⸻

Используемые технологии
 • Go
 • Apache Kafka
 • PostgreSQL
 • Docker
 • Docker Compose

⸻

Структура проекта

event-driven-order-system
│
├ cmd
│   ├ api
│   │   └ main.go
│   │
│   └ worker
│       └ main.go
│
├ internal
│   ├ handler
│   ├ kafka
│   ├ model
│   ├ repository
│   └ service
│
├ docker
│   └ docker-compose.yml
│
├ go.mod
└ README.md


⸻

Запуск проекта

1. Запуск инфраструктуры

Запустить Kafka, Zookeeper и PostgreSQL:

docker compose up -d


⸻

2. Запуск API сервиса

go run cmd/api/main.go

API будет доступен по адресу:

http://localhost:8080


⸻

3. Запуск Worker

Worker читает сообщения из Kafka и обрабатывает их параллельно с использованием goroutines.

go run cmd/worker/main.go


⸻

Пример запроса

PowerShell:

Invoke-RestMethod `
-Uri http://localhost:8080/orders `
-Method POST `
-ContentType "application/json" `
-Body '{"user_id":1,"product":"iphone","price":1000}'


⸻

Пример обработки Worker

worker 3 processed order {UserID:1 Product:iphone Price:1000}


⸻

Возможности системы
 • событийно-ориентированная архитектура
 • Kafka producer и consumer
 • Worker pool на goroutines
 • хранение данных в PostgreSQL
 • контейнеризация через Docker

⸻

Возможные улучшения
 • Retry механизм
 • Dead Letter Queue
 • Idempotent consumer
 • Graceful shutdown
 • Метрики и мониторинг
:::
