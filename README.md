# 🚀 Go API Template — Gin + PostgreSQL + sqlc

![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go)
![Framework](https://img.shields.io/badge/Gin-v1.12-00ADD8?style=flat-square)
![Database](https://img.shields.io/badge/PostgreSQL-316192?style=flat-square&logo=postgresql&logoColor=white)
![SQL](https://img.shields.io/badge/sqlc-Type--Safe-blue?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

A production-ready Go project template for building RESTful APIs. Built with **Gin** as the HTTP framework, **PostgreSQL** as the database, and **sqlc** for type-safe SQL code generation. This template follows a clean **layered architecture** pattern that separates concerns and makes the codebase scalable, testable, and maintainable.

---

## ✨ Features

- **Layered Architecture** — Clean separation between Controller, Service, and Repository layers
- **Type-Safe SQL** — No ORM overhead; write raw SQL and let `sqlc` generate Go code
- **PostgreSQL Ready** — Pre-configured database connection with health check and timeout
- **Environment Config** — `.env` file support via `godotenv`
- **Scalable Structure** — Pre-organized folder structure ready for rapid feature development
- **Migration Support** — Dedicated directory for database migration files

---

## 🛠 Tech Stack

| Technology                                   | Purpose                  |
| -------------------------------------------- | ------------------------ |
| [Go 1.25+](https://go.dev/)                  | Programming language     |
| [Gin v1.12](https://gin-gonic.com/)          | HTTP web framework       |
| [PostgreSQL](https://www.postgresql.org/)    | Relational database      |
| [sqlc](https://sqlc.dev/)                    | SQL-to-Go code generator |
| [lib/pq](https://github.com/lib/pq)          | PostgreSQL driver for Go |
| [godotenv](https://github.com/joho/godotenv) | `.env` file loader       |

---

## 📁 Project Structure

```text
.
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── config/
│   │   └── database.go             # PostgreSQL connection setup
│   │
│   ├── controllers/                # HTTP request handlers
│   │   └── <feature>_controller.go
│   │
│   ├── dto/                        # Data Transfer Objects (request/response)
│   │   └── <feature>_dto.go
│   │
│   ├── middlewares/                 # HTTP middleware (auth, CORS, logging)
│   │   └── <middleware>.go
│   │
│   ├── models/                     # Domain/business models
│   │   └── <feature>_model.go
│   │
│   ├── routes/                     # Route group definitions
│   │   └── routes.go
│   │
│   └── service/                    # Business logic layer
│       └── <feature>_service.go
│
├── db/
│   ├── migrations/                 # SQL migration files (up/down)
│   │   ├── 000001_create_xxx.up.sql
│   │   └── 000001_create_xxx.down.sql
│   │
│   └── query/                      # sqlc query definitions
│       └── <feature>.sql
│
├── .env                            # Environment variables (not committed)
├── sqlc.yaml                       # sqlc configuration
├── go.mod                          # Go module definition
└── go.sum                          # Dependency checksums
```

### Directory Breakdown

#### `cmd/api/`

Application entry point. `main.go` is responsible for:

- Initializing the Gin HTTP server
- Loading environment variables from `.env`
- Establishing and verifying the database connection
- Registering routes and starting the server

#### `internal/`

All application-internal packages live here. The `internal/` directory is a Go convention that prevents external packages from importing these packages.

| Directory      | Responsibility                                                                                      |
| -------------- | --------------------------------------------------------------------------------------------------- |
| `config/`      | Infrastructure configuration (database connection, external services)                               |
| `controllers/` | HTTP handlers — receives requests, validates input, calls service layer, returns HTTP response      |
| `dto/`         | Data Transfer Objects — defines the shape of request payloads and response bodies                   |
| `middlewares/` | HTTP middlewares — authentication, authorization, CORS, request logging, rate limiting              |
| `models/`      | Domain models — represents business entities and maps database results to application-level structs |
| `routes/`      | Route definitions — groups endpoints by feature and applies middlewares                             |
| `service/`     | Business logic — orchestrates repository calls, applies business rules, handles transactions        |

#### `db/`

| Directory     | Responsibility                                                          |
| ------------- | ----------------------------------------------------------------------- |
| `migrations/` | SQL migration files for schema versioning (up = apply, down = rollback) |
| `query/`      | Raw SQL query files used by `sqlc` to generate type-safe Go code        |

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/<your-username>/template-golang.git
cd template-golang
```

### 2. Configure Environment Variables

Copy and edit the `.env` file:

```bash
cp .env.example .env
```

Fill in your database credentials:

```env
APP_PORT=8080
DATABASE_URL=postgres://<user>:<password>@localhost:5432/<db_name>?sslmode=disable
```

### 3. Create the Database

```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE your_db_name;
\q
```

### 4. Run Database Migrations

If using [golang-migrate](https://github.com/golang-migrate/migrate):

```bash
migrate -path db/migrations -database "$DATABASE_URL" up
```

### 5. Generate sqlc Code

```bash
sqlc generate
```

### 6. Install Dependencies

```bash
go mod tidy
```

### 7. Run the Server

```bash
go run ./cmd/api/main.go
```

The server starts at `http://localhost:8080`. Verify with:

```bash
curl http://localhost:8080/
# Response: Back IS Running
```

---

## 🔧 Development Workflow

Follow these steps to add a new feature (e.g., `products`):

### Step 1 — Write the Migration

```bash
# Create migration files
migrate create -ext sql -dir db/migrations -seq create_products_table
```

Edit the generated `.up.sql`:

```sql
CREATE TABLE IF NOT EXISTS products (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    price       DECIMAL(10,2) NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  TIMESTAMP
);
```

Apply:

```bash
migrate -path db/migrations -database "$DATABASE_URL" up
```

### Step 2 — Write sqlc Queries

Create `db/query/product.sql`:

```sql
-- name: CreateProduct :one
INSERT INTO products (name, price) VALUES ($1, $2)
RETURNING id, name, price, created_at, updated_at;

-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL;

-- name: ListProducts :many
SELECT * FROM products WHERE deleted_at IS NULL ORDER BY id LIMIT $1 OFFSET $2;
```

Generate:

```bash
sqlc generate
```

### Step 3 — Create DTO

Create `internal/dto/product_dto.go`:

```go
package dto

type CreateProductRequest struct {
    Name  string  `json:"name" binding:"required"`
    Price float64 `json:"price" binding:"required,gt=0"`
}

type ProductResponse struct {
    ID        int32   `json:"id"`
    Name      string  `json:"name"`
    Price     float64 `json:"price"`
    CreatedAt string  `json:"created_at"`
}
```

### Step 4 — Create Model

Create `internal/models/product_model.go` to map database results:

```go
package models

import "time"

type Product struct {
    ID        int32
    Name      string
    Price     float64
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Step 5 — Create Service

Create `internal/service/product_service.go`:

```go
package service

type ProductService struct {
    queries *db.Queries  // sqlc generated
}

func NewProductService(q *db.Queries) *ProductService {
    return &ProductService{queries: q}
}

func (s *ProductService) Create(ctx context.Context, req dto.CreateProductRequest) (*models.Product, error) {
    // Business logic + call sqlc query
}
```

### Step 6 — Create Controller

Create `internal/controllers/product_controller.go`:

```go
package controllers

type ProductController struct {
    service *service.ProductService
}

func NewProductController(s *service.ProductService) *ProductController {
    return &ProductController{service: s}
}

func (ctrl *ProductController) Create(c *gin.Context) {
    // Parse request → call service → return response
}
```

### Step 7 — Register Routes

Add routes in `internal/routes/routes.go`:

```go
productCtrl := controllers.NewProductController(productService)

api := router.Group("/api/v1")
{
    products := api.Group("/products")
    {
        products.POST("/", productCtrl.Create)
        products.GET("/", productCtrl.List)
        products.GET("/:id", productCtrl.GetByID)
    }
}
```

---

## 📌 Environment Variables

| Variable       | Description                     | Example                                                      |
| -------------- | ------------------------------- | ------------------------------------------------------------ |
| `APP_PORT`     | Port the HTTP server listens on | `8080`                                                       |
| `DATABASE_URL` | PostgreSQL connection string    | `postgres://user:pass@localhost:5432/dbname?sslmode=disable` |

---