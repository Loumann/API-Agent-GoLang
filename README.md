#  Agent REST API Golang



###  Клонирование репозитория



- ``git clone https://github.com/Loumann/Api.git``


###  Установка зависимостей

- ``go mod download``

### 4. Настройка базы данных - PostgreSQL 
- Создайте базу данных AgentAPI.
- Настройте параметры подключения в файле config.json.

   #### config.json
  {
          "host": "localhost",
          "port": "5432",
          "username": "",
          "password": "",
          "dbname": "AgentAPI",
          "ssl_mode": "disable"
}



Для создания записи используем тело JSON:
- .../agents  для тестовой выборкы

### Тело json
{
    "agentname": "",
    "status": ""
}


## Структура проекта

```plaintext
agent-rest-api/
├── cmd/
│   └── main.go
├── configs/
│   └── cfg.go
├── handler/
│   └── agent.go
│   └── handler.go
├── models/
│   └── structAgent.go
│   └── structQuest.go
├── repos/
│   └── database.go
├── go.mod
├── go.sum
