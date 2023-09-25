# Authentication
## `POST` api/unprotected/users/signup
use this to create a new user from a username, password and optionally a display name
```ts
type Body = {
    username: string,
    display?: string, // default: `username`, dashes replaced with spaces and put in title case
    password: string, // UNHASHED!
}
type Returns = User | Error
```

## `GET ` api/unprotected/users/login
logs in a user for 24 hours by setting the JWT cookie
```ts
type Body = {
    username: string,
    password: string, // UNHASHED!
}
type Returns = {
    success: string // just a success message, can be ignored.
} | Error
// Also sets a cookie on the client, `token`, which contains the JWT.
```