# Migrations

This directory contains **SQL migration files** for database schema versioning.

## Naming Convention

```
<sequence>_<description>.up.sql     # Apply migration
<sequence>_<description>.down.sql   # Rollback migration
```

## Tool

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations.

### Install

```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Go install
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Commands

```bash
# Create a new migration
migrate create -ext sql -dir db/migrations -seq create_products_table

# Apply all pending migrations
migrate -path db/migrations -database "$DATABASE_URL" up

# Rollback last migration
migrate -path db/migrations -database "$DATABASE_URL" down 1

# Rollback all migrations
migrate -path db/migrations -database "$DATABASE_URL" down

# Force a specific version (fix dirty state)
migrate -path db/migrations -database "$DATABASE_URL" force <version>

# Check current version
migrate -path db/migrations -database "$DATABASE_URL" version
```

## Example

**`000001_create_users_table.up.sql`**

```sql
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  TIMESTAMP
);
```

**`000001_create_users_table.down.sql`**

```sql
DROP TABLE IF EXISTS users;
```