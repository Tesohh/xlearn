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
}
```

## `GET ` api/user/me
returns the currently logged in user
```ts
type Returns = User | Error
```