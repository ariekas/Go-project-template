# Middlewares

This directory contains **HTTP middleware** functions that execute before or after request handlers.

## Naming Convention

```
<middleware_name>.go
```

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