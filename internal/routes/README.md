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
	"database/sql"

	"template-golang/internal/controllers"
	"template-golang/internal/middlewares"
	"template-golang/internal/service"

	db "template-golang/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

func New(conn *sql.DB) *gin.Engine {
	router := gin.Default()

	// Initialize sqlc queries
	queries := db.New(conn)

	// Initialize layers
	productService := service.NewProductService(queries)
	productCtrl := controllers.NewProductController(productService)

	// Public routes
	router.GET("/", func(c *gin.Context) {
		c.String(200, "API is running")
	})

	// API v1 group
	api := router.Group("/api/v1")
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

	return router
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
app.Run(":" + os.Getenv("APP_PORT"))
```
