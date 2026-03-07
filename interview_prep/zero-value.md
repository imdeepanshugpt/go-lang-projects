# Zero Value in Go

In Go, every variable is automatically initialized to a _zero value_ when it
is declared without an explicit initializer. The exact zero value depends on
the type:

| Type Category     | Example Type               | Zero Value        |
| ----------------- | -------------------------- | ----------------- |
| Numeric           | `int`, `float64`           | `0`, `0.0`        |
| Boolean           | `bool`                     | `false`           |
| String            | `string`                   | `""` (empty)      |
| Pointer           | `*T`                       | `nil`             |
| Slice/Map/Channel | `[]T`, `map[K]V`, `chan T` | `nil`             |
| Interface         | `interface{}`              | `nil`             |
| Function          | `func()`                   | `nil`             |
| Array             | `[N]T`                     | all elements zero |
| Struct            | `struct{...}`              | all fields zero   |

## Properties

- There is no uninitialized memory in Go; every variable has a well-defined
  value after declaration.
- The zero value is the result of declaring a variable with the `var` keyword
  and no initializer:
  ```go
  var x int        // x == 0
  var s string     // s == ""
  var p *int       // p == nil
  ```
- When using `new(T)`, Go allocates zeroed storage and returns a pointer to it.
- Composite types (arrays, structs) have their elements/fields set to the
  zero value of their element type.
- The zero value is useful in APIs that rely on "default" configuration
  without needing explicit initialization.

## Example

```go
package main

import "fmt"

func main() {
    var a int               // zero value 0
    var b bool              // false
    var c []string          // nil slice
    var m map[string]int    // nil map
    var p *float64          // nil pointer

    fmt.Printf("%v %v %v %v %v\n", a, b, c, m, p)
}
```

```text
0 false [] map[] <nil>
```

> Note: printing a nil slice or map shows its type but no elements; attempting
> to read or write to a nil map or slice may cause a run-time panic (except
> reading length/capacity of a slice is safe).

---

Having predictable zero values simplifies many Go patterns and reduces the
need for constructors just to set initial defaults.
