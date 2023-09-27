# Org
## Definition
```ts
type Org = {
    name: string,
    tag: string,
    secret: string,
}
```

## `POST` api/org/new
```ts
type Body = Org
type Returns = Org | Error
```