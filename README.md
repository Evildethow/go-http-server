# Go HTTP Server

Example HTTP Server written in Golang.

# Prerequisites (optional)

* Enable [Dependabot](https://docs.github.com/en/code-security/getting-started/dependabot-quickstart-guide)
* Label Pull Requests with either `enhancement` or `bug` for release notes

# Usage

## Build

```shell
make
```

## Run

HTTP server runs on port `8080` by default.

Endpoints:

* `/ping`

```shell
make run
```

## Docker Run

```shell
make docker-run
```