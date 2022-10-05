# graphql-vite
GraphQL subscriptions client in ReactJS + Vite

### Requirements

- [go 1.18](https://go.dev/doc/install)
- [gqlgen](https://github.com/99designs/gqlgen)
- [curl](https://help.ubidots.com/en/articles/2165289-learn-how-to-install-run-curl-on-windows-macosx-linux)
- [nvm](https://github.com/nvm-sh/nvm#install--update-script)
- [pnpm](https://pnpm.io/installation)
- [jq](https://github.com/stedolan/jq/wiki/Installation)

### Create User

```graphql
mutation createUser {
    createUser(input: {username: "u", password: "p"}) {
        id
        username
        createdAt
        updatedAt
    }
}
```

### Create Todo


```graphql
mutation createTodo {
    createTodo(input: {text: "todo", done: false, userID: 1}) {
        id
        done
        updatedAt
        createdAt
    }
}
```
