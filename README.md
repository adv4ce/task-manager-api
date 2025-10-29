# 🧩 TaskManagerAPI

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Status](https://img.shields.io/badge/Status-Active-success)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen)

Простое и расширяемое REST API для управления задачами, написанное на **Go (Golang)** с использованием фреймворка **Gin**.  
Проект демонстрирует архитектуру на основе слоёв (`cmd`, `internal`, `handlers`, `models`, `services`) и аккуратное разделение логики.

---

## 📖 Содержание

- [Описание](#description)
- [Функциональность](#features)
- [Структура проекта](#project-structure)
- [Технологии](#technologies)
- [Установка и запуск](#installation)
- [Примеры API-запросов](#api-examples)
- [Переменные окружения](#environment-variables)
- [Тестирование](#testing)

---

## 🧠 Описание {#description}

**TaskManagerAPI** — это REST API-сервис для работы с задачами.  
Он позволяет создавать, получать, обновлять и удалять задачи через HTTP-запросы.  
Архитектура проекта построена по принципу **чистой структуры**, что упрощает масштабирование и добавление новых функций.

---

## ⚙️ Функциональность {#features}

- Проверка состояния API (`/health`)
- Создание задачи
- Получение задачи по ID
- Обновление задачи
- Удаление задачи
- Хранение данных в памяти (через структуру `TaskLib`)
- Обработка ошибок и унифицированные ответы (`APIResponse`)
- Тесты для основных обработчиков

---

## 🗂️ Структура проекта {#project-structure}

```
TaskManagerAPI/
├── cmd/
│   └── server/
│       └── main.go           			 # Точка входа, инициализация сервера
├── internal/
│   ├── config/
│   │   ├── config.go         				# Загрузка конфигурации
│   │   └── config.yml.example        # Параметры окружения
│   ├── handlers/
│   │   ├── api.go            # Реализация HTTP-эндпоинтов
│   │   ├── router.go         # Настройка маршрутов Gin
│   │   └── handlers_test.go  # Тесты API
│   ├── models/
│   │   └── models.go         # Модели данных
│   └── services/
│       ├── services.go       # Бизнес-логика и операции с задачами
│       └── inputFormat.go    # Форматирование входных данных
├── go.mod / go.sum           # Зависимости
├── .env.example              # Переменные окружения
├── .gitignore								# Игнорируемые файлы
└── README.md                 # Документация
```

---

## 🧰 Технологии {#technologies}

- **Go** — язык программирования
- **Gin** — HTTP-фреймворк для API
- **Go modules** — управление зависимостями
- **YAML / ENV** — конфигурация
- **Testing (Go test)** — модульные тесты

---

## 🚀 Установка и запуск {#installation}

1. **Клонировать репозиторий:**

   ```bash
   git clone https://github.com/yourusername/TaskManagerAPI.git
   cd TaskManagerAPI/TaskManagerAPI
   ```

2. **Установить зависимости:**

   ```bash
   go mod tidy
   ```

3. **Настройка конфигурации:**
```bash
# Вариант A: через Environment Variables
cp .env.example .env
# Отредактируйте .env при необходимости

# Вариант B: через YAML
cp internal/config/config.yml.example internal/config/config.yml
# Отредактируйте config.yml при необходимости
```

4. **Запустить сервер:**

   ```bash
   go run cmd/server/main.go
   ```

5. **Проверить работу API:**
   ```bash
   curl http://localhost:8080/health
   ```
   **Ответ:**
   ```json
   {
   	"status": "healthy"
   }
   ```

---

## 📡 Примеры API-запросов {#api-examples}

### ✅ Проверка состояния

`GET /api/health`

**Ответ:**

```json
{
	"status": "healthy"
}
```

---

### ➕ Создать задачу

`POST /api/tasks`

**Тело запроса:**

```json
{
	"title": "Закончить проект",
	"description": "Реализовать API и написать README",
	"priority": "high"
}
```

**Ответ:**

```json
{
	"status": "successful",
	"message": "Task created",
	"data": {
		"title": "Закончить проект",
		"description": "Реализовать API и написать README",
		"priority": "high"
	}
}
```

---

### 🔍 Получить задачу

`GET /api/tasks/:id`

**Пример:**

```bash
curl http://localhost:8080/api/tasks/1
```

**Ответ:**

```json
{
	"status": "successful",
	"data": {
		"id": 1,
		"title": "Закончить проект",
		"description": "Реализовать API и написать README",
		"priority": "high"
	}
}
```

---

### ✏️ Обновить задачу

`PUT /api/tasks/:id`

**Тело запроса:**

```json
{
	"title": "Обновлённая задача",
	"description": "Добавлены тесты и обработка ошибок",
	"priority": "medium"
}
```

**Ответ:**

```json
{
	"status": "successful",
	"message": "Task updated"
}
```

---

### ❌ Удалить задачу

`DELETE /api/tasks/:id`

**Ответ:**

```json
{
	"status": "successful",
	"message": "Task deleted"
}
```

---

## ⚙️ Переменные окружения {#environment-variables}

| Переменная | Описание                                      | Пример  |
| ---------- | --------------------------------------------- | ------- |
| `PORT`     | Порт, на котором запускается сервер           | `8080`  |
| `MODE`     | Режим работы приложения (`debug` / `release`) | `debug` |

---

## 🧪 Тестирование {#testing}

```bash
go test -v ./internal/handlers/

go test -cover ./internal/handlers/
```

⭐️ **Если вам понравился мой проект, не забудьте поставить звёздочку на GitHub!**
