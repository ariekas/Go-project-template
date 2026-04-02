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
```