# ternary

[![Build Status](https://travis-ci.org/otiai10/ternary.svg?branch=master)](https://travis-ci.org/otiai10/ternary)

Ternary expression for golang, to enjoy code-golf :golf:

```go
port := ternary.String(os.Getenv("PORT"))("8080")
// Returns "8080" if "PORT" == ""
```

# why

If you want

```go
status := 500
if flag {
    status = 200
}
```

Write

```go
status := ternary.If(flag).Int(200, 500)
```

If you want

```go
res := map[string]interface{}{
    "user":       user,
    "expired_at": nil,
}
if user.Expire() {
    res["expire_at"] = time.Now()
}
json.NewEncoder(wr).Encode(res)
```

Write

```go
json.NewEncoder(wr).Encode(map[string]interface{}{
    "user":       user,
    "expired_at": ternary.If(user.Expire()).Interface(time.Now(), nil),
})
```

Enjoy :golf: !
