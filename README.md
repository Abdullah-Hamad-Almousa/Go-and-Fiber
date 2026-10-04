# day2

This repo is not useful. It is a scratch pad of Go snippets with no app, no library, and no problem to solve. `master` is whatever `main.go` happens to do right now. The other branches are placeholders: names and notes only. They are not created yet.

## master

`main.go` is one file that runs three tiny demos and prints a line of 37 dashes between them.

The pointer helpers are still in the file but do not run. `zeroval` takes an `int` by value and sets that copy to `0`, so the caller would not change. `zeroptr` takes an `*int` and sets the pointed-at value to `0`. The demo that would call them, print the address, and try `new(42)` is commented out. `new(42)` would not compile anyway, because `new` wants a type, not a value.

What actually runs:

1. A `Player` starts at 100 health. `TakeDamage` is a pointer method, so `TakeDamage(30)` changes the real struct. It prints `Hero Health: 70`.
2. A `Counter` starts at 0. `Increment` is also a pointer method. It is called twice, then the unexported `count` is printed as `2` (same package, so `main` can see it).
3. A hardcoded JSON blob `{"name": "Abdullah", "age": 27}` is decoded with `encoding/json/v2` into a `User` (`Name`, `Age`, json tags). It prints `Abdullah is 27`, encodes the struct back to JSON, and writes `user.json`. Marshal and write errors are ignored.

The dash loops use `for x := range 37`. The `x++` inside does nothing to the iteration count. Each loop just prints 37 `-` characters and a newline.

## placeholder branches

Fourteen branches that do not exist yet. Each line is a name and what that branch is for.

| branch | what I did |
| --- | --- |
| `master` | Go Setup & Syntax *T and &T. |
| `day2` | for loop and range. make(), append(), delete(). |