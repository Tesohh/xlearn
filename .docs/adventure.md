# Adventures
## Definitions
```ts
type Adventure = {
	name: string
	tag: string
	description: string
	steps: string[]
}
type Step = {
	name: string
	tag: string
	description: string
	content: string // WARNING: will extremely likely change in type in the near future 
	xpAward: number
	coinsAward: number   
	energyCost: number   

	children: string[] // the tags to children
}
```

## `GET ` api/org/@{orgtag}/adventure/all
Returns a list of adventures in the current org (steps are omitted to save resources)
```ts
type Returns = Adventure[] | null | Error
```


## `GET ` api/org/@{orgtag}/adventure/@{advtag}
Returns a single adventure from a tag
```ts
type Returns = Adventure | Error
```



## `POST ` api/org/@{orgtag}/adventure/new
Modifiers: `admin`
```ts
type Body = {
	name: string,
	description: string,
}
```
```ts
type Returns = Adventure | Error
```