# go-optional

The `optional` type for golang, implemented with generics (go 1.18).

## TLDR
```golang
var v optional.Type[int] = optional.New(123)
if unwrapped, ok := v.Value(); ok {
    unwrapped // 123
}

v.IsNil()      // false
v.ForceValue() // 123
```

## Methods

### initializer
- `optional.New(XX)`
- `optional.Nil()`
- (`optional.FromPtr(XXX)` returns optional.Nil() if pointer is nil) 

### value
- `Value()`
```golang
var v = optional.New(123)
if unwrapped, ok := v.Value(); ok {
    unwrapped // 123
}
```
- `ForceValue()`
- `IsNil()`


- `ValueOrDefault(XX)`
- `ValueOrLazyDefault(f)`

### transform

- Map

```golang
optional.Map( optional.New(123), func(t int) string {
    return fmt.Sprintf("hello %d", t)
}) // optional.Type[string]
```

- Compact

```golang
var v = optional.New(optional.New(123))
optional.Compact(v) // optional.Type[int]
```

### serialization
- json.Marshal and json.Umarshal


# License

MIT