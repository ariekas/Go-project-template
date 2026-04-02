# Query

This directory contains **SQL query files** used by [sqlc](https://sqlc.dev/) to generate type-safe Go code.

## Naming Convention

```
<feature>.sql
```

## sqlc Annotations

Each query **must** have a `-- name:` comment annotation:

| Annotation    | Go Return Type             | Use Case                         |
| ------------- | -------------------------- | -------------------------------- |
| `:one`        | Single struct + `error`    | Get by ID, Create with RETURNING |
| `:many`       | Slice of structs + `error` | List queries                     |
| `:exec`       | `error` only               | Delete, Update without RETURNING |
| `:execresult` | `sql.Result` + `error`     | When you need affected row count |

## Example

**`product.sql`**

```sql
-- name: CreateProduct :one
INSERT INTO products (name, price)
VALUES ($1, $2)
RETURNING id, name, price, created_at, updated_at;

-- name: GetProductByID :one
SELECT id, name, price, created_at, updated_at
FROM products
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListProducts :many
SELECT id, name, price, created_at, updated_at
FROM products
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET name = $2, price = $3, updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, name, price, created_at, updated_at;

-- name: SoftDeleteProduct :exec
UPDATE products
SET deleted_at = NOW()
WHERE id = $1;
```

## Generate

After writing or updating query files, run:

```bash
sqlc generate
```

Generated Go code will be placed in `internal/db/sqlc/` (as configured in `sqlc.yaml`).

## Rules

- Use parameterized queries (`$1`, `$2`, ...) — never concatenate user input into SQL
- Prefer `RETURNING` clause to get the inserted/updated row in a single query
- Use soft delete (`deleted_at IS NULL`) instead of hard delete when appropriate
- Keep queries **focused** — one operation per query function
