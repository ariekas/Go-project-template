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

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(s *service.ProductService) *ProductController {
	return &ProductController{service: s}
}

func (ctrl *ProductController) Create(c fiber.Ctx) error {
	var req dto.CreateProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := ctrl.service.Create(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": result})
}
```
