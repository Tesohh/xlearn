# Authentication

## `POST` api/user/signup

Modifiers: `unprotected`
use this to create a new user from a username, password and optionally a display name

```ts
type Body = {
  username: string;
  display?: string; // default: `username`, dashes replaced with spaces and put in title case
  password: string; // UNHASHED!
};
type Returns = User | Error;
```

## `POST` api/user/login

Modifiers: `unprotected`
logs in a user for 24 hours by setting the JWT cookie

```ts
type Body = {
  username: string;
  password: string; // UNHASHED!
};
type Returns =
  | {
      success: string; // just a success message, can be ignored.
    }
  | Error;
// Also sets a cookie on the client, `token`, which contains the JWT.
```

## `GET ` api/user/logout

removes the jwt cookie from the user

```ts
type Returns =
  | {
      success: string; // just a success message, can be ignored.
    }
  | Error;
```

## `POST` api/user/recover

Modifiers: `unprotected`
changes the user's password to the new one provided.
if the pin is incorrect, RecoverAttempts will be incremented.
if it reaches 3 or more, the account is locked from recovering and must ask an admin to reset the counter.

Note: this doesn't login the user, so it must be done manually with /login

```ts
type Body = {
  username: string;
  pin: string;
  new_password: string; // UNHASHED!
};
type Returns = Success | Error;
```
