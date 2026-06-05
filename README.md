# Monolithic Modular Go

A Clean Architecture boilerplate using Go, Echo, and SQLite.

## Architecture

```
domain/           → entity + interfaces (ports)
usecase/          → implements domain usecase interface
handler/          → HTTP adapter (outer)
repository/       → database adapter (outer)
```

- **Dependency Rule**: all dependencies point inward to `domain/`
- **Interface**: defined in `domain/`, implemented by outer layers
- **Mapping**: happens at adapter boundaries (handler ↔ usecase, repository ↔ database)
- **Validation**: at handler level, not usecase (usecase = pure business logic)
- **Error Handling**: `*apperror.AppError` returned from usecase, handled by middleware
- **Logging**: structured logging with trace ID via `context.Context`

## Structure

```
cmd/
├── server/          # all modules (users, products)
├── inventory/      # inventory preset (products, product_groups)
└── migrate/        # migration CLI

internal/
├── bootstrap/       # server bootstrap
├── config/         # config loader (yaml + env)
├── driver/         # db drivers (sqlite)
├── middleware/     # trace_id, access_log, error_handler
├── migration/      # Go-based migrations (001_users.go, etc)
├── modules/
│   ├── users/
│   │   ├── domain/user/
│   │   ├── usecase/user/
│   │   ├── handler/user/
│   │   │   └── model/
│   │   └── repository/user/
│   │       └── model/
│   └── products/
│       ├── domain/product/
│       ├── domain/product_group/
│       ├── usecase/product/
│       ├── usecase/product_group/
│       ├── handler/product/
│       ├── handler/product_group/
│       ├── repository/product/
│       └── repository/product_group/
├── router/         # core router
└── utils/

pkg/
├── response/         # OK, Created, OKMessage helpers
├── apperror/         # error codes + AppError struct
└── logger/          # structured logging with trace ID

docs/
├── docs.go          # swagger package
├── swagger.json     # generated swagger doc
└── swagger.yaml     # generated swagger doc
```

## Quick Start

```bash
# copy env
cp .env.example .env

# run migrations
make migrate

# run all modules
make run

# run inventory preset
go run cmd/inventory/main.go

# build
make build
make build-inventory
```

## Configuration

`config.yml` — server & module selection

```yaml
server:
  addr: ":8080"
  api_prefix: "/api/v1"
  log_level: "info"  # debug, info, warn, error

modules:
  - users
  - products
```

`.env` — database

```bash
DB_DRIVER=sqlite
DB_DSN=app.db
```

### Presets

| Config | Preset | Modules |
|--------|--------|---------|
| `config.yml` | `all` (default) | users, products |
| `config.inventory.yml` | `inventory` | products, product_groups |

## Migration

```bash
# migrate up (run all migrations)
make migrate

# migrate down (rollback last migration)
make migrate-down

# check migration status
make migrate-status

# drop all tables
make migrate-drop
```

Migration files are Go files in `internal/migration/`:
- `001_users.go`
- `002_products.go`
- `003_product_groups.go`

Each file registers `Up` and `Down` functions via `Register()` in `init()`.

## Swagger

```bash
# generate swagger docs
make swag
```

Swagger UI available at `/swagger/index.html` when server is running.

Add annotations to handler functions:
```go
// Create godoc
//
//	@Summary		Create user
//	@Description	Create new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.CreateUserRequest	true	"User data"
//	@Success		201		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Router			/api/v1/users [post]
func (h *UserHandler) Create(c echo.Context) error {
```

## API Endpoints

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/users` | Create user |
| GET | `/api/v1/users` | List users |
| GET | `/api/v1/users/:id` | Get user |

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/products` | Create product |
| GET | `/api/v1/products` | List products |
| GET | `/api/v1/products/:id` | Get product |

### Product Groups

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/product-groups` | Create group |
| GET | `/api/v1/product-groups` | List groups |
| GET | `/api/v1/product-groups/:id` | Get group |

## Response Format

### Success

```json
{
  "success": true,
  "data": { ... }
}
```

```json
{
  "success": true,
  "message": "deleted successfully"
}
```

### Error

```json
{
  "success": false,
  "code": "NOT_FOUND",
  "message": "user not found"
}
```

## Docker

```bash
# all modules
docker compose up server

# inventory preset
docker compose up inventory
```

## Build Per Module

```bash
# all modules → ~15MB
go build -o bin/server cmd/server/main.go

# inventory only → ~12MB
go build -o bin/inventory cmd/inventory/main.go
```