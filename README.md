# DO NOT USE THIS

YOU HAVE BETTER review your database schema and client rquirements of API response.

I don't want to write such fuckin' code any more.

```go
response := map[stirng]interface{}{
	"expire": func() interface{} {
		// because client requires "expire" field as "null"
		if user.Expire.IsZero() {
			return nil
		}
		return user.Expire
	}(),
}
```

to be

```go
response := map[stirng]interface{}{
    "expire": ternary.Returns((user.Expire.IsZero()), nil, user.Expire),
}
```

# つらい
