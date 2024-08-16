#  Agent REST API Golang



###  Клонирование репозитория



- ``git clone https://github.com/Loumann/Api.git``


###  Установка зависимостей

- ``go mod download``

###  Настройка базы данных - PostgreSQL 
- Создайте базу данных AgentAPI.
- Настройте параметры подключения в файле config.json.

   #### config/cfg.json
  {
          "host": "",
          "port": "",
          "dbname": "",
          "ssl_mode": ""
}
  #### .env.local
POSTGRES_USER=""
POSTGRES_PASSWORD=""



Для создания записи используем тело JSON:
- :8081/agents  для тестовой выборкы

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
