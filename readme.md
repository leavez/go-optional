# go-optional

An optional type for golang with generics

```golang

v := optional.New(123)
if unwrapped, ok := v.Value(); ok {
	print(unwrapped == 123)
}

print(v.ForceValue() == 123)

v2 := optional.Nil[int]()
print(v2.IsNil() == true)


```

# License

MIT