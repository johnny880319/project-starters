# My Routine For Starting a TypeScript Project

All commands should be run from the `typescript/` directory. Adjust paths as needed if you change the structure.

## Initialize a project

Install [unzip](https://itsfoss.com/unzip-linux/) if you do not have it already.

```bash
sudo apt update && sudo apt install unzip -y
```

Install [fnm](https://fnm.vercel.app/) if you do not have it already.

```bash
curl -fsSL https://fnm.vercel.app/install | bash
```

Install [pnpm](https://pnpm.io/installation) if you do not have it already.

```bash
fnm install 24
fnm default 24
corepack enable pnpm
```

Initialize a project with current settings:

```bash
pnpm install
```

## Manage dependencies

Manage dependencies:

```bash
pnpm add <package-name>
pnpm remove <package-name>
```

Manage development dependencies (excluded from production):

```bash
pnpm add -D <package-name>
pnpm remove <package-name>
```

Use `pnpm install` to install dependencies from `package.json` and update `pnpm-lock.yaml` (for local dev). Add `--frozen-lockfile` to strictly install from the lockfile without updating it (for CI/CD).

```bash
pnpm install
pnpm i --frozen-lockfile
```

## Linting, type checking and formatting

TypeScript has built-in type checking. We use [Biome](https://biomejs.dev/) as the linter and formatter:

You may need to run the following command after cloning the repo to set up Biome:

```bash
pnpm biome migrate --write
```

To run these tools from the command line:

```bash
pnpm biome check --write <file-or-dir>
pnpm biome format --write <file-or-dir>
pnpm exec tsc --noEmit
```

Lint config lives in `biome.json`.

## Project structure

This project uses the standard `src` layout.

* **src/**: application source code
* **tests/**: contains all unit and integration tests
* **./**: configuration files and documentation

## Testing

This template uses [Vitest](https://vitest.dev/) for unit tests.

**Naming conventions:** Vitest automatically discovers tests. Ensure your test files are named `*.test.ts` (or `*.spec.ts`) and your test functions start with `test` or `it`.

To run the tests:

```bash
pnpm vitest run
```

## TODO: Dockerfile (Podman/Docker)

To build and run the Docker image:

```bash
podman build -t localhost/my-typescript-app .
podman run --rm localhost/my-typescript-app
```

Dockerfile template lives in `Dockerfile` and `.dockerignore`.

## VS Code integration

Install the **Biome** extension for linting and formatting support in VS Code.

VS Code settings lives in `../.vscode/settings.json`.
