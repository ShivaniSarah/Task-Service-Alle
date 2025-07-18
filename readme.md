# Task Management Microservice (Go)

This project implements a simple **Task Management System** in Go using microservice architecture principles. It supports CRUD operations, pagination, and filtering by task status.

---

## Features

- Create, Read, Update, Delete (CRUD) tasks
- Pagination on `GET /tasks`
- Filtering by task `status`
- Clean separation of concerns
- Designed for easy scalability and future extensibility

---

## Tech Stack

- **Language**: Go
- **Framework**: Gin (HTTP server)
- **ORM**: GORM (with SQLite)
- **Dependency Management**: Go Modules

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

---

##  📁 Complete Project Structure


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
git clone  https://github.com/ShivaniSarah/Task-Service-Alle.git

cd task-service

# Install dependencies
go mod tidy

# Run server
go run main.go

# Server will be available at http://localhost:8080
```

---

## Design Decisions

### 1. Microservices Principles Applied

- **Single Responsibility**: Handlers handle HTTP, services handle logic, repositories handle DB.
- **Loose Coupling**: Interfaces are used to decouple layers.
- **Scalable**: Ready for horizontal scaling via Docker/Kubernetes.

### 2. Pagination and Filtering

- Supported via query params:  
  `GET /tasks?status=Completed&limit=10&offset=0`
- Efficient DB-level filtering and paging

### 3. Extensibility

- Easy to add a **User Service** (e.g., for task ownership)
- Interfaces allow easy switch of DBs or transport layers

### 4. Scalability

- The service is stateless and suitable for horizontal scaling.

- Replace local SQLite with a shared DB like PostgreSQL or MySQL.

- Run multiple instances using Docker or Kubernetes.

- Use a load balancer (e.g., NGINX, cloud LB) to distribute traffic.

- Kubernetes can manage replicas, autoscaling, and failovers.

- Add a /healthz endpoint for readiness and liveness checks.

- Centralized logging via ELK/Grafana is recommended.

- Easily extendable to microservice architecture with gRPC or message queues.


---

## Inter-Service Communication (Future Scope)

| Option     | Use Case                                     |
|------------|----------------------------------------------|
| REST       | For human-readable APIs, external services   |
| gRPC       | For internal high-performance communication  |
| Kafka/NATS | For async updates/events like task completed |

---

## Author

**Shivani Agrawal [ShivaniSarah@github]**

---



## Commands Used

```bash
brew install sqlite
brew install go
go mod init task-service
go mod tidy
go run main.go
```

### Generate Mocks

```bash
go install github.com/vektra/mockery/v2@latest
export PATH=$PATH:$(go env GOPATH)/bin
mockery --name=TaskService --dir=service --output=service/mocks --outpkg=mocks
```

### Run Tests

```bash
go get github.com/stretchr/testify
go test ./handler
go test ./...
```

---

## SQLite Commands

```bash
sqlite3 tasks.db
.tables
SELECT * FROM tasks;
.quit
```

---


## API Documentation


### 1. Create a Task

**POST** `http://localhost:8080/tasks`

**Request:**

```json
{
  "title": "Finish Go Assignment",
  "description": "Pending"
}
```

**Response:**

```json
{
  "id": 28,
  "title": "Task last",
  "description": "My last task",
  "status": "CREATED"
}
```

---

### 2. Get All Tasks

**GET** `http://localhost:8080/tasks?status=Pending&limit=5&offset=0`

**Response:**

```json
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
```

---

### 3. Get Task by ID

**GET** `http://localhost:8080/tasks/1`

**Response:**

```json
{
  "id": 26,
  "title": "Finish Go Assignment Updated",
  "description": "Changed the description",
  "status": "MODIFIED"
}
```

---

### 4. Update a Task

**PATCH** `http://localhost:8080/tasks/1`

#### Modify Title

```json
{
  "title": "Finish Go Assignment Updated"
}
```

**Response:**

```json
{
  "id": 26,
  "title": "Finish Go Assignment more",
  "description": "sample",
  "status": "MODIFIED"
}
```

#### Mark as Completed

```json
{
  "title": "Finish Go Assignment Updated",
  "status": "COMPLETED"
}
```

**Response:**

```json
{
  "id": 26,
  "title": "Finish Go Assignment more",
  "description": "sample",
  "status": "COMPLETED"
}
```

---

### 5. Delete a Task

**DELETE** `http://localhost:8080/tasks/1`


## Drawbacks abd Future Enhancements:

1. Pagination not has hasNext , pageISize, pageNumber, totalRecords

2. DTO and entity class

3. Middleware in golang instead of aspect (custom annotations in java)

4. If it is marked as Completed , don’t mark it completed multiple  times

5. Validation for update not done


---



