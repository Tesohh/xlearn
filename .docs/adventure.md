# Adventures
## Definitions
```ts
type Adventure = {
	name: string
	tag: string
	description: string
	steps: string[]
}
```

## `GET ` api/org/@{orgtag}/adventure/all
Modifiers: `protectorg`
Returns a list of adventures in the current org (steps are omitted to save resources)
```ts
type Returns = Adventure[] | null | Error
```


## `GET ` api/org/@{orgtag}/adventure/@{advtag}
Modifiers: `protectorg`
Returns a single adventure from a tag
```ts
type Returns = Adventure | Error
```

## `POST` api/org/@{orgtag}/adventure/@{advtag}
Modifiers: `protectorg`
can edit an adventure. Note: if you put in a tag it will be ignored as it can never be changed.
Modifiers: `admin`
```ts
type Body = Adventure
type Returns = Success | Error
```


## `POST ` api/org/@{orgtag}/adventure/new
Modifiers: `admin`, `protectorg`
```ts
type Body = {
	name: string,
	description: string,
}
```
```ts
type Returns = Adventure | Error
```