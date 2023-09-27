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