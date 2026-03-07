# Error Handling in Go

Go takes a simple, explicit approach to errors rather than using exceptions. The
language encourages returning error values and checking them immediately. This
section summarizes the idioms and patterns you’ll encounter.

## The `error` type

An `error` is a built-in interface type:

```go
package builtin

type error interface {
    Error() string
}
```

Any type implementing `Error() string` satisfies the interface.

### Creating errors

```go
import "errors"

func f() error {
    return errors.New("something went wrong")
}

// or with fmt
return fmt.Errorf("failed to open %s: %w", path, err)
```

The `%w` verb wraps an existing error for later unwrapping.

## Returning errors

Most functions that may fail return an `error` as the last result:

```go
func ReadFile(name string) ([]byte, error)
```

Callers handle the error immediately:

```go
data, err := ReadFile("foo.txt")
if err != nil {
    // handle/return
    log.Fatal(err)
}
// use data
```

This “if err != nil” check is the idiomatic form.

## Custom error types

You can define your own error struct when you need more context:

```go
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string {
    return fmt.Sprintf("%s %s: %v", e.Op, e.Path, e.Err)
}
```

Standard library errors such as `os.PathError`, `net.OpError`, etc. use this pattern.

## Error wrapping and unwrapping

Since Go 1.13, errors can wrap others using `fmt.Errorf` with `%w` or
`errors.Wrap` in third-party libraries.

```go
if err := do(); err != nil {
    return fmt.Errorf("do failed: %w", err)
}
```

To examine wrapped errors:

```go
if errors.Is(err, os.ErrNotExist) {
    // matched
}
if errors.As(err, &pathErr) {
    // extracted *os.PathError
}
```

- `errors.Is` checks for equality along the chain.
- `errors.As` finds the first error in the chain that can be assigned to
  the target type.

## Panics and recover

Panic is reserved for truly unexpected conditions; they stop normal flow.
`defer`/`recover` can turn a panic into an error for controlled recovery:

```go
func safe() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    // ... code that may panic
    return nil
}
```

Panics are not idiomatic for ordinary error reporting.

## Conventions & tips

- Name the error return value `err` by convention.
- Add context before propagating: `return fmt.Errorf("read %s: %w", name, err)`
- Avoid ignoring errors—explicit `_ =` or comment `// ignore error` if you do.
- For APIs that never fail, omit the error return.
- Do not use exceptions; Go has none.

> ⚠️ Some packages return multiple errors or use `[]error`; still check them.

---

Go’s explicit error checks may feel verbose at first, but they make control
flow clear and easy to reason about. By returning errors as values, Go avoids
hidden control transfer and keeps error handling part of the type system.
