# Org
## Definition
```ts
type Org = {
    name: string,
    tag: string,
    secret: string,
    adventures: Adventure[]
}
```

## `POST` api/org/new
```ts
type Body = {
    name: string,
    secret: string,
}
type Returns = Org | Error
```