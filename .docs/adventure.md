# Adventures
## Definitions
```ts
type Adventure = {
	name: string
	tag: string
	description: string
	steps: Step[]
}
type Step = {
	name: string
	tag: string
	description: string
	content: string // WARNING: will extremely likely change in type in the near future 
	xpAward: number
	coinsAward: number   
	energyCost: number   

	children: Step[]
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



## `POST ` api/admin/org/@{orgtag}/adventure/new
```ts
type Body = {
	name: string,
	description: string,
}
```
```ts
type Returns = Adventure | Error
```