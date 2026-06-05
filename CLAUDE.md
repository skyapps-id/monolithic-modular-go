# CLAUDE.md

Clean Architecture boilerplate: Go + Echo + SQLite (`database/sql`, no ORM). Module path: `github.com/skyapps-id/monolithic-modular-go`.

## Layers (per module)
`domain/` entity + interfaces (ports) → `usecase/` business logic → `handler/` HTTP adapter + `repository/` DB adapter.
Dependencies point inward to `domain/`. Adapters depend on domain interfaces, never vice versa.

## Conventions
- One file per use case/handler/repo method (`create.go`, `find_by_id.go`, `find_all.go`); constructor + struct in `usecase.go`/`handler.go`/`repository.go`.
- Mapping at boundaries: handler `model/` ↔ domain (`ToInput`, `*ToResponse`); repository `model/` ↔ domain (`*FromEntity`).
- Validation in handler, NOT usecase. Usecase = pure logic.
- Errors: usecase/handler return `*apperror.AppError` (`apperror.New`, `NotFoundError`, `ValidationError`); `middleware.ErrorHandler` renders them.
- Logging: `logger.{Debug,Info,Error}Ctx(ctx, "module.layer.action.event", "key", val)` — trace ID flows via `context.Context`.
- Responses: `response.OK` / `response.Created` / `response.OKMessage`.
- A module implements `router.Module` (`Name()`, `RegisterRoutes(*echo.Group)`); wired in `NewModule(db)`. Register new modules in `cmd/server/main.go`.

## Layout
- `cmd/{server,inventory,migrate}/main.go` — entrypoints / presets.
- `internal/modules/{users,products}/` — feature modules (vertical slices).
- `internal/{bootstrap,config,driver,middleware,migration,router}/` — infra.
- `pkg/{apperror,logger,response}/` — shared, no business logic.
- `docs/` — generated swagger (`make swag`).

## Commands
- `make run` — server (all modules) · `make build` / `make build-inventory`
- `make migrate` / `migrate-down` / `migrate-status` / `migrate-drop`
- `make test` · `make lint` (golangci-lint) · `make tidy` · `make swag`
- Module selection via `config.yml` `modules:` list; preset config: `config.inventory.yml`. DB via env `DB_DRIVER`/`DB_DSN`.

## Migrations
Go files in `internal/migration/` (`NNN_name.go`), each calls `Register(Up, Down)` in `init()`. Add new file with next number.

## When adding a feature/endpoint
1. domain (entity, repository + usecase interface) → 2. usecase impl → 3. repository impl → 4. handler + handler/model → 5. route in module `RegisterRoutes` → 6. swagger annotation + `make swag`.
Match the existing per-file split; keep the dependency direction.
