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