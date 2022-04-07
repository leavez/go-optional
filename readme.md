# go-optional

An optional type for golang with generics

```golang
var v optional.Type[int] = optional.New(123)
if unwrapped, ok := v.Value(); ok {
    unwrapped // 123
}

v.IsNil()      // false
v.ForceValue() // 123

v2 := optional.Nil[int]()
v2.IsNil() // true

optional.Map(v, func(t int) string {
    return fmt.Sprintf("hello %d", t)
}) // Type[string]

v3 := optional.New(optional.New(123))
optional.Compact(v3) // Type[int]
```

# License

MIT