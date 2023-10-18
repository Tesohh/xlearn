# User
## Definition
```ts
type User = {
    display: string,
    username: string,
    passwordhash: string,
    xp: number,
    level: number,
    coins: number,
    role: 0 | 1 | 2, // user, teacher, admin
    joined_orgs: string[]
}
```

## `GET ` api/user/me
returns the currently logged in user
```ts
type Returns = User | Error
```

## `POST` api/user/joinorg/{code}
joins an org given a code. Don't need to specify what org to join!
```ts
type Returns = {
    joined: string // tag of the org i joined
} | Error
```