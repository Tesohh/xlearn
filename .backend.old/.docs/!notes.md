# Important notes
## Errors
```ts
type Error = {
    error: string
}
```

## Tags
Think about tags as the primary keys of almost everything.
I'm not using id's as it would be cumbersome to work with in Sveltekit (trust me i've been there)
Tags are non editable: Please don't try to edit them.
In case you mistakenly edit a tag, in 99% of cases i will have placed a safeguard that disallows it.

## Modifiers
* `admin`: current user needs admin role
* `teacher`: current user needs teacher role
* `unprotected`: can be accessed without being logged in
* `protectorg`: current user needs to be in the org