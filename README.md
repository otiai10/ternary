# ternary

Ternary expression for golang.

I just want to write DRY code... :(

# why

- code golf

```go
// want
status := "Waiting"
if flag {
    status = "Ready"
}

// write
status := ternary.If(flag).String("Ready", "Waiting")
```

- value or `nil`

```go
// when client requires "or null" field
res := map[string]interface{}{
    "name":   user.Name,
    "expire": func() interace{} {
        if user.Expire.IsZero() {
            return nil
        }
        return user.Expire
    },
}

// write
res := map[string]interface{}{
    "name":   user.Name,
    "expire": ternary.Returns(user.Expire.IsZero(), nil, user.Expire),
}
```