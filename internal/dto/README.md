# DTO (Data Transfer Objects)

This directory contains **request and response structures** used for data transfer between the client and the API.

## Naming Convention

```
<feature>_dto.go
```

## Example

```go
package dto

// --- Request DTOs ---

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

// --- Response DTOs ---

type ProductResponse struct {
	ID        int32   `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// --- Pagination ---

type PaginationRequest struct {
	Page  int32 `form:"page" binding:"min=1"`
	Limit int32 `form:"limit" binding:"min=1,max=100"`
}
```

## Rules

- Request DTOs use `binding` tags for Gin validation
- Response DTOs use `json` tags for JSON serialization
- Keep DTOs separate from database models — they define the **API contract**, not the database schema
