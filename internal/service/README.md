# Service

This directory contains the **business logic layer**. Services orchestrate data flow between controllers and the database (via sqlc-generated queries).

## Naming Convention

```
<feature>_service.go
```

## Example

```go
package service

import (
	"context"

	"template-golang/internal/dto"
	"template-golang/internal/models"

	db "template-golang/internal/db/sqlc"
)

type ProductService struct {
	queries db.Querier // sqlc generated interface
}

func NewProductService(q db.Querier) *ProductService {
	return &ProductService{queries: q}
}

func (s *ProductService) Create(ctx context.Context, req dto.CreateProductRequest) (*models.Product, error) {
	row, err := s.queries.CreateProduct(ctx, db.CreateProductParams{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &models.Product{
		ID:        row.ID,
		Name:      row.Name,
		Price:     row.Price,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}

func (s *ProductService) GetByID(ctx context.Context, id int32) (*models.Product, error) {
	row, err := s.queries.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Product{
		ID:        row.ID,
		Name:      row.Name,
		Price:     row.Price,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}

func (s *ProductService) List(ctx context.Context, req dto.PaginationRequest) ([]models.Product, error) {
	rows, err := s.queries.ListProducts(ctx, db.ListProductsParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	})
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for _, row := range rows {
		products = append(products, models.Product{
			ID:        row.ID,
			Name:      row.Name,
			Price:     row.Price,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}

	return products, nil
}
```
