# Task Management Microservice (Go)

This project implements a simple **Task Management System** in Go using microservice architecture principles. It supports CRUD operations, pagination, and filtering by task status.

---

## Features

* Create, Read, Update, Delete (CRUD) tasks
* Pagination on `GET /tasks`
* Filtering by task `status`
* Clean separation of concerns
* Designed for easy scalability and future extensibility

---

## Tech Stack

* Language: **Go**
* Framework: **Gin** (HTTP server)
* ORM: **GORM** (with SQLite)
* Dependency Management: **Go Modules**

---

## 📁 Project Structure

```
.
├── main.go              # Entry point
├── handler/             # HTTP layer (Gin handlers)
├── service/             # Business logic
├── repository/          # DB operations
├── model/               # Data model (integrated in repository)
├── go.mod / go.sum      # Dependency management
```

## Complete Project Structure

```
task-service/
├── go.mod
├── go.sum                  # generated after `go mod tidy`
├── main.go
├── handler/
│   └── task_handler.go
├── repository/
│   ├── task.go
│   └── task_repository.go
├── service/
│   └── task_service.go
```

---

## Running the Service

```bash
# Clone repository
$ git clone https://github.com/pvnptl/task-service.git
$ cd task-service

# Run
$ go run main.go

# Server runs at http://localhost:8080
```

---

##  API Documentation

### Create Task

`POST /tasks`

```json
{
  "title": "Write Go service",
  "status": "Pending"
}
```

### Get All Tasks (with pagination & filtering)

`GET /tasks?status=Completed&limit=5&offset=10`

Response:

```json
[
  { "id": 3, "title": "sample title", "description": "test description" },
  ...
]
```

### Get Task by ID

`GET /tasks/:id`

### Update Task

`PATCH /tasks/:id`

```json
{
  "title": "Write unit tests",
  "status": "COMPLETED"
}
```

### Delete Task

`DELETE /tasks/:id`

---

## Design Decisions

### 1. **Microservices Principles Applied**

* **Single Responsibility**: Handlers only handle HTTP, Services handle logic, Repositories handle DB.
* **Loose Coupling**: Interfaces are used to decouple layers.
* **Scalable**: Easy to scale horizontally by containerizing (e.g., Docker + Kubernetes).

### 2. **Pagination and Filtering**

* Query params: `GET /tasks?status=Completed&limit=10&offset=0`
* Implemented at DB query level (efficient for large datasets)

### 3. **Extensibility**

* Easy to add a **User Service** in future (e.g., for task ownership)
* Interface-based design allows switching databases or transport protocols

---

## Inter-Service Communication (Future Scope)

If adding a **User Service**, here are options for communication:

| Option     | Use Case                                     |
| ---------- | -------------------------------------------- |
| REST       | For human-readable APIs, external services   |
| gRPC       | For internal high-performance communication  |
| Kafka/NATS | For async updates/events like task completed |

---

## Author

Shivani Agrawal

---


## Commands ran

 brew install sqlite
 brew install go
 go mod tidy
 go run main.go




go install github.com/vektra/mockery/v2@latest
go env GOPATH
export PATH=$PATH:$(go env GOPATH)/bin
mockery --name=TaskService --dir=service --output=service/mocks --outpkg=mocks


 go get github.com/stretchr/testify
 go test ./handler
 go test ./...

## Commands ran for sqlite3

sqlite3 tasks.db
.tables
SELECT * FROM tasks;
.quit


## Manual Testing Steps

1. Create a Task

POST http://localhost:8080/tasks

Body (JSON):

{
  "title": "Finish Go Assignment",
  "description": "Pending"
}

Response

{
    "id": 28,
    "title": "Task last",
    "description": "My last task",
    "status": "CREATED"
}

2. Get All Tasks

GET http://localhost:8080/tasks

Query Params (optional):

status=Pending

pageSize=5

page=1

Example:

http://localhost:8080/tasks?status=Pending&limit=5&offset=0

Response 

[
    {
        "id": 3,
        "title": "Task 3",
        "description": null,
        "status": "Created"
    },
    {
        "id": 4,
        "title": "Task 4",
        "description": null,
        "status": "Created"
    }
]

3. Get Task by ID

GET http://localhost:8080/tasks/1

Response

{
    "id": 26,
    "title": "Finish Go Assignment Updated",
    "description": "Changed the description",
    "status": "MODIFIED"
}

4. Update a Task

PATCH http://localhost:8080/tasks/1

To modify:

Body (JSON):
{
  "title": "Finish Go Assignment Updated",
  
}

Response

{
    "id": 26,
    "title": "Finish Go Assignment more",
    "description": "sample",
    "status": "MODIFIED"
}

To mark as completed:

PATCH http://localhost:8080/tasks/1

Body (JSON):
{
  "title": "Finish Go Assignment Updated",
  "status": "COMPLETED"
}

Response

{
    "id": 26,
    "title": "Finish Go Assignment more",
    "description": "sample",
    "status": "COMPLETED"
}

5. Delete a Task

DELETE http://localhost:8080/tasks/1