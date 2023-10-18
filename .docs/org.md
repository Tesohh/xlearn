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
Modifiers: `protectorg`
```ts
type Returns = Org | Error
```

## `POST` api/org/@{orgtag}
can edit an org. Note: if you put in a tag it will be ignored as it can never be changed.
Modifiers: `admin`, `protectorg`
```ts
type Body = Org
type Returns = Success | Error
```

## `GET ` api/org/@{orgtag}/meta
returns an org, but omits the adventures for performance reasons (returns only the metadata)
Modifiers: `protectorg`
```ts
type Returns = Omit<Org, "adventures"> | Error
```

## `GET ` api/org/@{orgtag}/code/{uses}
generates a code to join the org, with `uses` uses
if you put an invalid value in uses, it will default to 1.
Modifiers: `protectorg`
```ts
type Returns = {
    code: string
    uses: int
} | Error
```