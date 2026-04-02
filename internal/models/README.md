# Models

This directory contains **domain/business models** that represent application-level entities.

## Naming Convention

```
<feature>_model.go
```

## Purpose

Models act as a **mapping layer** between raw database results (sqlc-generated structs) and the application's response structures. They hold the business-level representation of data.

## Example

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

## Mapping Function Example

Convert sqlc-generated database struct to application model:

```go
package models

import db "template-golang/internal/db/sqlc"

func ToProductModel(row db.Product) Product {
	return Product{
		ID:        row.ID,
		Name:      row.Name,
		Price:     row.Price,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func ToProductModels(rows []db.Product) []Product {
	products := make([]Product, len(rows))
	for i, row := range rows {
		products[i] = ToProductModel(row)
	}
	return products
}
```

## Rules

- Models **must not** depend on HTTP-specific packages (`gin`, `net/http`)
- Models are used in the **service layer** to return clean data to controllers
- Mapping functions live alongside the model they produce
