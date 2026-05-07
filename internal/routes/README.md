# Routes

This directory contains **route group definitions** that map HTTP endpoints to their corresponding controllers and middlewares.

## Naming Convention

```
routes.go          # Main route setup
<feature>_routes.go  # Feature-specific route groups (optional)
```

## Example

```go
package routes

import (
	"template-golang/internal/controllers"
	"template-golang/internal/middlewares"
	"template-golang/internal/service"

	db "template-golang/internal/db/sqlc"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(pool *pgxpool.Pool) *fiber.App {
	app := fiber.New()

	// Initialize sqlc queries
	queries := db.New(pool)

	// Initialize layers
	productService := service.NewProductService(queries)
	productCtrl := controllers.NewProductController(productService)

	// Public routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "API is running",
		})
	})

	// API v1 group
	api := app.Group("/api/v1")
	{
		// Product routes
		products := api.Group("/products")
		{
			products.GET("/", productCtrl.List)
			products.GET("/:id", productCtrl.GetByID)
			products.POST("/", productCtrl.Create)
			products.PUT("/:id", productCtrl.Update)
			products.DELETE("/:id", productCtrl.Delete)
		}
	}

	// Protected routes (example)
	protected := api.Group("/admin")
	protected.Use(middlewares.AuthMiddleware())
	{
		// Admin-only endpoints here
	}

	return app
}
```

## Usage in `main.go`

```go
db, err := config.ConnectionDB()
if err != nil {
    log.Fatal(err)
}
defer db.Close()

app := routes.New(db)
app.Listen(":" + os.Getenv("APP_PORT"))
```
