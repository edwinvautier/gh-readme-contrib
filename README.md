# gh-readme-contrib

[![Build](https://github.com/edwinvautier/gh-readme-contrib/actions/workflows/ci.yml/badge.svg)](https://github.com/edwinvautier/gh-readme-contrib/actions/workflows/ci.yml)

## Weekly activity

<div align="center">
  <img src="https://gh-readme-contrib.herokuapp.com/api/edwinvautier/gh-readme-contrib"/>
</div>

## Add chart to your repository

Your repository must be public, or at least you have to deploy the app with a custom github API token.
To add the image on your repository, simply add the following line to your README: 

```md
![](https://gh-readme-contrib.herokuapp.com/api/author/repository)
```

## Set config colors

The following query parameters are available:

| name           | description                                                                        | type | example                                                    |
| -------------- | ---------------------------------------------------------------------------------- | ---- | ---------------------------------------------------------- |
| main           | the main color (chart curve)                                                       | hex  | `/api/edwinvautier/gh-readme-contrib?main=f0f`             |
| bg             | the background color                                                               | hex  | `/api/edwinvautier/gh-readme-contrib?bg=f0f`               |
| text           | text color                                                                         | hex  | `/api/edwinvautier/gh-readme-contrib?text=f0f`             |
| display_number | say if you want to display the current week commits number or not, default to true | bool | `/api/edwinvautier/gh-readme-contrib?display_numbre=false` |

## Setup

> If you want more informations about the available make commands, run `make help`

If you already have the go.mod and go.sum files you can run :

```sh
make init
# or
make start
```

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
