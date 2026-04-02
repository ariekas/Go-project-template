# Controllers

This directory contains **HTTP request handlers** (controllers). Each controller is responsible for:

- Parsing and validating incoming HTTP requests
- Calling the appropriate service method
- Returning HTTP responses with proper status codes

## Naming Convention

```
<feature>_controller.go
```

## Example

```go
package controllers

import (
	"net/http"

	"template-golang/internal/dto"
	"template-golang/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(s *service.ProductService) *ProductController {
	return &ProductController{service: s}
}

func (ctrl *ProductController) Create(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}
```

## Rules

- Controllers **must not** contain business logic — delegate to the service layer
- Controllers **must not** access the database directly
- Use DTOs for request binding and response formatting
