# My Routine For Starting a Go Project

## Create a module

From the `go/` directory, initialize the module (replace with your real repository path):

```bash
go mod init github.com/<github-username>/<project-name>/go
```

## Manage dependencies

```bash
# Add a dependency
go get <package-path>

# Clean and sync go.mod/go.sum
go mod tidy
```

## Linting, type checking, and formatting

Go type checking is built into `go build` / `go test`. For linting, this template uses [golangci-lint](https://github.com/golangci/golangci-lint).

Install `golangci-lint` (pin a version for consistent CI results):

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.10.1
```

Run checks:

```bash
golangci-lint run
```

Lint config lives in `go/.golangci.yml`.

### VS Code integration

Install the **Go** extension (`golang.go`), then add this to `.vscode/settings.json`:

```json
{
  "[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": "explicit"
    }
  },
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package"
}
```

## Project structure

This template follows the common Go layout:

- `cmd/`: application entrypoints (`main` packages)
- `internal/`: private code, importable only within this module
- `pkg/`: reusable library code for external usage (optional)

```
go/
├── cmd/
│   └── my-app/
│       ├── main.go
│       └── main_test.go
├── internal/
│   └── my-module/
│       ├── module.go
│       └── module_test.go
├── pkg/
│   └── my-package/
│       ├── package.go
│       └── package_test.go
├── .gitignore
├── .golangci.yml
├── README.md
├── go.mod
└── go.sum
```

## Testing

Use [Testify](https://github.com/stretchr/testify) for assertions and mocking.

```bash
go get github.com/stretchr/testify
```

Go discovers tests in `*_test.go` files. Test functions should use the `TestXxx` naming pattern.

```bash
go test ./...
go test -v ./...
go test -cover ./...
```

## Dockerfile (Podman/Docker)

This template includes `go/Dockerfile` and `go/.dockerignore`.
Run these commands from the `go/` directory:

```bash
podman build -t localhost/my-go-app .

# Run
podman run --rm localhost/my-go-app
```

Notes:
- This starter app is CLI-style, so `EXPOSE` is intentionally not enabled by default.
- Add `EXPOSE <port>` only when your app actually listens on that port.

## Live reloading (optional)

Use [Air](https://github.com/air-verse/air) for local live reload:

```bash
go install github.com/air-verse/air@latest
air init
air
```
