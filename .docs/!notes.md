# Important notes
## Modifiers
Modifiers are the first thing after /api/.
Some routes have no modifiers.

### `unprotected`
By default, all routes need to have a JWT containing the current username.
If a route has the `unprotected` modifier, it can be accessed freely by anyone, even with no account

### `admin`
The route can only be accessed by users with the Admin role.