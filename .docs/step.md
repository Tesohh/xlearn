# Step
## Definition
```ts
type Step = {
    name: string;
    tag: string;
    description: string;
    content: string; // markdown is stored here
    category: "lesson" | "exercise" | "project";
    xp_award: number;
    coins_award: number;
    energy_cost: number;

    children: string[][]; // slice of slices of tags to other Steps
}
```

## Tree structure
> if you cant preview this, look it up on github or download a mermaid previewer extension

Suppose you have an adventure that looks like this
```mermaid
graph TD
  S1(Step1) --> S2(Step2)
	S2 --> S3(Python) & S4(JS)

	S3 --> str(Strings) --> ints(Integers)

	S4 --> console(Console)
	
	ints & console --> S5(Step5)
```
in (pseudo)code would be:
```go
// Adventure children:
[
    {"step1"},
    {
        "step2", 
        children: [
            [
                "python", "strings", "integers"
            ],
            [
                "js", "console"
            ]
        ]
    }
    {"step5"}
]
```