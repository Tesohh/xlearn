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

## api/org/@{tag}
```ts
type Returns = Org | Error
```

## api/org/@{tag}/meta
returns an org, but omits the adventures for performance reasons (returns only the metadata)
```ts
type Returns = Omit<Org, "adventures"> | Error
```