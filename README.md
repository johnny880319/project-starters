# My routine of starting a new project

## Initialize git repository

```bash
git init
```

Set up project-specific Git identity:

```bash
git config --local user.name "Your Name"
git config --local user.email "your.email@example.com"
```

## Create essential files

```bash
touch README.md Makefile .gitignore
```

### VS Code integration

```bash
mkdir -p .vscode && touch .vscode/settings.json
```

## First commit and push

Stage and commit locally:

```bash
git add .
git commit -m "chore: initial commit"
```

Create an empty repository on GitHub first, then push your initial commit:

```bash
git remote add origin git@github.com:your-account/<your-project-name>.git
git branch -M main
git push -u origin main
```

## Set up Git hooks

Install Lefthook if you haven't already:

```bash
go install github.com/evilmartians/lefthook@latest
```

Set up Git hooks with Lefthook:

```bash
lefthook install
```

Hook config lives in `lefthook.yml`.

### VS Code integration
Install **GitLens** and **Git Graph** extensions for advanced Git UI.

## Conventional workflow

Follow these industry standards for collaboration and versioning:

1. [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow)
2. [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
3. [Semantic Versioning](https://semver.org/)
4. [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
