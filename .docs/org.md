# Org
## Definition
```ts
type Org = {
    name: string,
    tag: string,
    secret: string,
    adventures: string[]
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

## `GET ` api/org/@{orgtag}
```ts
type Returns = Org | Error
```

## `POST` api/org/@{orgtag}
can edit an org. Note: if you put in a tag it will be ignored as it can never be changed.
```ts
type Body = Org
type Returns = Success | Error
```

## `GET ` api/org/@{orgtag}/meta
returns an org, but omits the adventures for performance reasons (returns only the metadata)
```ts
type Returns = Omit<Org, "adventures"> | Error
```