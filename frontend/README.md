# simple-todo frontend

this is the frontend project of simple-todo

## generate types and hooks

we use [@graphql-codegen](https://www.the-guild.dev/graphql/codegen) to generate the necessary `types` for typescript and the `hooks` for react.

follow the steps to add `@graphql-codegen` to the project

### steps:

- install the dev dependencies:

```console
pnpm install @graphql-codegen/cli @graphql-codegen/typescript @graphql-codegen/typescript-operations @graphql-codegen/typescript-urql @graphql-codegen/urql-introspection graphql-tag --save-dev
```

- create a file `codegen.yml`:

```yml
overwrite: true
schema: "http://localhost:5000/query"
documents: "src/**/*.graphql"
generates:
  src/graphql/generated.ts:
    plugins:
      - "typescript"
      - "typescript-operations"
      - "urql-introspection"
      - "typescript-urql"
```

- add generate script in your `package.json`:

```json
   "scripts":{
      "generate": "graphql-codegen --config codegen.yml"
   }
```

- run the generate script

```console
   pnpm generate
```

the script will search for any file `.graphql` in `src/**/*` folders and generate the needed types and hooks, make sure the backend is running in order for the `graphql-codegen` to find the `schema`.
