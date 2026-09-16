# Repository Guidelines

## Project Structure & Module Organization

This repository is a Go backend organized under `app/server`:

- `cmd/main.go` is the application entry point and wires middleware, configuration, MongoDB, and routes.
- `internal/config` contains environment and database configuration.
- `internal/module/auth` contains authentication routes, controllers, and services.
- `internal/module/user` contains user-related models and module code.
- Tests should be placed beside the package they cover in files named `*_test.go`.

Keep business logic in services, HTTP concerns in controllers, and route registration in module-level files.

## Build, Test, and Development Commands

Run commands from the repository root:

```text
go run ./app/server/cmd     # Run the API locally
go build ./...              # Compile all packages
go test ./...               # Run all tests
gofmt -w path/to/file.go    # Format changed Go files
```

The `.env` file supplies local configuration such as `PORT` and database settings. Do not commit secrets. The existing `make dev` and `make prod` targets are intended for environment-specific runs; verify their entry-point path before using them.

## Coding Style & Naming Conventions

Use standard `gofmt` formatting and idiomatic Go naming: exported identifiers use PascalCase, unexported identifiers use camelCase, and package names are short and lowercase. Prefer clear constructor names such as `NewService`. Return errors to callers and keep handlers responsible for HTTP status and response formatting.

## Testing Guidelines

Use Go's standard `testing` package. Name tests `TestTypeOrFunction_Scenario`, keep unit tests close to the implementation, and run `go test ./...` before submitting changes. Add coverage for authentication failures, malformed requests, and successful responses when modifying login behavior.

## Commit & Pull Request Guidelines

Recent commits use short Conventional Commit-style subjects, for example `feat: implement authentication module` and `refactor: simplify database connection`. Follow the same `type: imperative summary` pattern (`feat`, `fix`, `refactor`, `test`, or `docs`).

Pull requests should explain the change, identify affected endpoints or packages, mention configuration changes, and include test results. Include request/response examples when changing API behavior and link the relevant issue when one exists.

## Security & Configuration

Never commit tokens, SSH keys, AWS credentials, or Discord webhooks. Configure secrets through GitHub repository or environment settings, and use least-privilege permissions for workflow jobs.

## Permissions

Global rule:

- Ask the user first before making any code change.
- Show the intended change for review when possible.
- Wait for the user to accept or reject the change before editing project code.
- Always keep the changes in the code to be as simple as possible, straight to the point, with no comments added

### Allow

- Read tracked project files needed for the task.
- Read source code under `src/`, configuration under `config/`, and docs such as `README.md`, `CLAUDE.md`, and this file.
- Create new source or documentation files when they are required for the requested change.
- Edit application code, route files, models, middleware, utilities, tests, and markdown documentation.
- Update `package.json` when the task explicitly requires script or dependency changes.
- Run safe repo-local commands such as `rg`, `ls`, `sed`, `git status`, `pnpm lint`, and `pnpm test`.

### Deny

- Do not read, print, or copy secrets from `.env`, `.env.dev`, or any credential file.
- Do not modify `.env`, `.env.dev`, or other secret-bearing files unless the user explicitly asks.
- Do not modify `node_modules/`, generated caches, or log files.
- Do not change `pnpm-lock.yaml` unless dependency work is part of the task.
- Do not delete files, rename major directories, or rewrite large parts of the codebase without explicit approval.
- Do not run destructive git or shell commands such as `git reset --hard`, `git checkout --`, or broad `rm` operations.
- Do not alter deployment/infrastructure files (`Dockerfile`, `docker-compose.yml`, `ecosystem.config.cjs`) unless the task explicitly requires it.
