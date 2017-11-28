# Bytes

bytes is a package to perform common transformations between go primitive types to bytes array.

Its semantics are simple, just use `[Primitive type][From/To]Bytes`

For example, to convert from **int** to **[]byte** use:

```go
var n int = 155
bytAr := bytes.IntToBytes(n)
fmt.Println(bytAr)
```

And to convert the array it back to **int**:

```go
m := bytes.IntFromBytes(bytar)

if n != m {
    panic("Not equal")
}
```

It also provides a convenient `Equal` function to compare two bytes arrays by checking element by element. And a `ToBytes` function to infer the type of the incoming interface