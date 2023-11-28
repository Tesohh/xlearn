# User

## Definition

```ts
type User = {
  display: string;
  username: string;
  passwordhash: string;
  xp: number;
  level: number;
  coins: number;
  role: 0 | 1 | 2; // user, teacher, admin
  joined_orgs: string[];
  settings: Settings;
};

type Settings = {
  language: string;
  theme: string;
}
```

## `GET ` api/user/me
returns the currently logged in user

```ts
type Returns = User | Error;
```

## `GET ` api/user/me/settings/edit
```ts
type Body = Settings; // if a field is left empty it won't be modified
type Returns = Settings | Error;
```

## `POST` api/user/org/leave/{code}
joins an org given a code. Don't need to specify what org to join!

```ts
type Returns =
  | {
      joined: string; // tag of the org i joined
    }
  | Error;
```

## `POST` api/user/org/leave/@{orgtag}
```ts
type Returns = {
    left: string // tag of the org i left
} | Error
```

## `POST` api/user/org/joined/
Returns an array of orgs the user joined + unprotected orgs. 
```ts
type Returns = Org[] | Error
```

## `POST` api/user/org/joined/tags
Returns an array of orgs the user joined + unprotected orgs, but only their tags.
```ts
type Returns = string[] | Error
```
