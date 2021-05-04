# gh-readme-contrib

[![CodeQL](https://github.com/edwinvautier/gh-readme-contrib/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/edwinvautier/gh-readme-contrib/actions/workflows/codeql-analysis.yml)
[![Build](https://github.com/edwinvautier/gh-readme-contrib/actions/workflows/ci.yml/badge.svg)](https://github.com/edwinvautier/gh-readme-contrib/actions/workflows/ci.yml)

![img](http://localhost:8000/api/edwinvautier/go-gadgeto)

## Setup

First time you run the project ? Run :

```sh
make first-run
```

This command creates the go mod and tidies dependencies and then call the make init command.

> If you want more informations about the available make commands, run `make help`

If you already have the go.mod and go.sum files you can run :

```sh
make init
# or
make start
```

The difference between make init and make start commands is that make init also copy .env.dist to .env and generates new RSA key.

## Branch naming convention

You branch should have a name that reflects it's purpose.

It should use the same guidelines as [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md) (`feat`, `fix`, `build`, `perf`, `docs`), followed by an underscore (`_`) and a very quick summary of the subject in **kebab case**.

Example: `feat_add-image-tag-database-relation`.

## Pull requests (PR)

Pull requests in this project follow two conventions, you will need to use the templates available in the [ISSUE_TEMPLATE](.github/ISSUE_TEMPLATE) folder :

-   Adding a new feature should use the [FEATURE_REQUEST](.github/ISSUE_TEMPLATE/FEATURE_REQUEST.md) template.
-   Reporting a bug should use the [BUG_REPORT](.github/ISSUE_TEMPLATE/bug_report.md) template.

If your pull request is still work in progress, please add "WIP: " (Work In Progress) in front of the title, therefor you inform the maintainers that your work is not done, and we can't merge it.

The naming of the PR should follow the same rules as the [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md)

## Linter

We use go linter [gofmt](https://blog.golang.org/gofmt) to automatically formats the source code.

you can run `make format` to auto-format your files.
