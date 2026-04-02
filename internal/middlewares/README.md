# Middlewares

This directory contains **HTTP middleware** functions that execute before or after request handlers.

## Naming Convention

```
<middleware_name>.go
```

## Common Middlewares

| Middleware      | Purpose                                     |
| --------------- | ------------------------------------------- |
| `auth.go`       | JWT/token authentication and authorization  |
| `cors.go`       | Cross-Origin Resource Sharing configuration |
| `logger.go`     | Request/response logging                    |
| `rate_limit.go` | API rate limiting                           |
| `recovery.go`   | Panic recovery                              |

## Example

```go
package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token here...

		c.Set("user_id", userID)
		c.Next()
	}
}
```

## Usage in Routes

```go
api := router.Group("/api/v1")
api.Use(middlewares.AuthMiddleware())
{
    api.GET("/profile", ctrl.GetProfile)
}
```

## Rules

- Middlewares **must** call `c.Next()` to pass control to the next handler, or `c.Abort()` to stop the chain
- Use `c.Set()` / `c.Get()` to pass data between middleware and handlers
