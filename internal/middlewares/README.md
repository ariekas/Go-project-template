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

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token here...

		c.Locals("user_id", userID)
		return c.Next()
	}
}
```
