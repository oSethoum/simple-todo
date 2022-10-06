# simple-todo
GraphQL subscriptions client in ReactJS + Vite + urql

## Requirements

- [go 1.18](https://go.dev/doc/install)
- [gqlgen](https://github.com/99designs/gqlgen)
- [curl](https://help.ubidots.com/en/articles/2165289-learn-how-to-install-run-curl-on-windows-macosx-linux)
- [nvm](https://github.com/nvm-sh/nvm#install--update-script)
- [pnpm](https://pnpm.io/installation)
- [jq](https://github.com/stedolan/jq/wiki/Installation)

## Help

```bash
make help
```
```text
Usage: make COMMAND
Commands :
help              - List available targets
deps              - Install dependencies
clean-frontend    - Cleanup frontend
install-frontend  - Install frontend
generate-frontend - Generate frontend
build-frontend    - Build frontend
run-frontend      - Run frontend
update-frontend   - Update frontend
clean-backend     - Cleanup backend
generate-backend  - Generate backend GraphQL source code
test-backend      - Run backend tests
build-backend     - Build backend GraphQL API
run-backend       - Run backend GraphQL API
get-backend       - Download and install go backend packages
update-backend    - Update backend dependencies to latest versions

```

## Run

### Terminal 1 - Start backend

```shell
make run-backend
```

Now you can open GraphiQL an in-browser tool for writing, validating, and testing GraphQL queries at [http://localhost:5000/playground](http://localhost:5000/playground)

### Terminal 2 - Start frontend

```shell
make run-frontend
```

Now you can onpe fronted UI at [http://localhost:5173/](http://localhost:5173/)
