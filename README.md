# Product Catalog API

REST API для управления каталогом товаров.

Проект выполнен в рамках тестового задания на позицию Junior Backend Developer.

## Технологии

* Go
* Gin
* GORM
* SQLite

## Функциональность

API поддерживает следующие операции:

* Создание товара
* Получение товара по ID
* Получение списка товаров
* Обновление товара
* Удаление товара
* Фильтрацию товаров по категории
* Фильтрацию по диапазону цен
* Пагинацию
* Валидацию входных данных
* Обработку ошибок
* Unit-тесты бизнес-логики

## Структура проекта

```text
product-catalog-api/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── apperror/
│   │   └── errors.go
│   │
│   ├── database/
│   │   └── database.go
│   │
│   ├── handler/
│   │   ├── error_response.go
│   │   └── product_handler.go
│   │
│   ├── model/
│   │   └── product.go
│   │
│   ├── repository/
│   │   └── product_repository.go
│   │
│   └── service/
│       ├── product_service.go
│       └── product_service_test.go
│
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Запуск проекта

### 1. Клонирование репозитория

```bash
git clone <repository-url>
```

### 2. Переход в папку проекта

```bash
cd product-catalog-api
```

### 3. Установка зависимостей

```bash
go mod download
```

### 4. Запуск приложения

```bash
go run ./cmd
```

После запуска сервер будет доступен по адресу:

```text
http://localhost:8080
```

База данных SQLite создаётся автоматически при запуске приложения.

## Health Check

```http
GET /health
```

Ответ:

```json
{
  "status": "ok"
}
```

---

# API

## Создание товара

```http
POST /products
```

Пример запроса:

```json
{
  "name": "iPhone 17",
  "description": "Новый смартфон",
  "price": 99999.99,
  "category": "electronics"
}
```

Пример ответа:

```json
{
  "id": 1,
  "name": "iPhone 17",
  "description": "Новый смартфон",
  "price": 99999.99,
  "category": "electronics",
  "created_at": "2026-09-03T12:00:00Z"
}
```

Статус ответа:

```text
201 Created
```

---

## Получение списка товаров

```http
GET /products
```

Пример:

```text
GET /products
```

### Query-параметры

| Параметр   | Описание                |
| ---------- | ----------------------- |
| limit      | Количество товаров      |
| offset     | Смещение                |
| category   | Фильтрация по категории |
| price_from | Минимальная цена        |
| price_to   | Максимальная цена       |

Пример запроса:

```text
GET /products?category=electronics&price_from=1000&price_to=200000&limit=10&offset=0
```

---

## Получение товара по ID

```http
GET /products/:id
```

Пример:

```text
GET /products/1
```

---

## Обновление товара

```http
PUT /products/:id
```

Пример запроса:

```json
{
  "name": "iPhone 17 Pro",
  "description": "Обновлённое описание",
  "price": 129999.99,
  "category": "electronics"
}
```

---

## Удаление товара

```http
DELETE /products/:id
```

Пример:

```text
DELETE /products/1
```

При успешном удалении:

```text
204 No Content
```

---

# Валидация

При создании или обновлении товара проверяются:

* Название товара не должно быть пустым
* Цена должна быть больше 0
* Категория не должна быть пустой

При фильтрации:

* `price_from` не может быть отрицательным
* `price_to` не может быть отрицательным
* `price_from` не может быть больше `price_to`
* `limit` должен быть больше 0
* `offset` не может быть отрицательным

---

# Ошибки

API использует стандартные HTTP-статусы:

| Статус | Описание                  |
| ------ | ------------------------- |
| 200    | Успешный запрос           |
| 201    | Товар успешно создан      |
| 204    | Товар успешно удалён      |
| 400    | Некорректный запрос       |
| 404    | Товар не найден           |
| 500    | Внутренняя ошибка сервера |

Пример ошибки:

```json
{
  "error": "product not found"
}
```

---

# Тестирование

Для запуска всех тестов:

```bash
go test ./...
```

Для подробного вывода:

```bash
go test -v ./...
```

Для форматирования кода:

```bash
go fmt ./...
```
