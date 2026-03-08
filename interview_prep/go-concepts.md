# Go Concepts: Panic/Recover and Receivers

## Panic vs Recover

### **panic**

- A built-in function used when the program encounters a serious error it can't handle.
- When called, it stops normal execution immediately, unwinds the stack, and begins to run deferred functions.
- If not handled, the program crashes and prints a stack trace.
- **Typical use**: errors that shouldn't happen in normal operation (e.g., invalid state, out-of-bounds slice access in library code).

```go
func doSomething() {
    panic("unexpected state")
}
```

### **recover**

- A built-in function that regains control of a panicking goroutine.
- Must be called within a deferred function; otherwise it returns `nil`.
- If `recover` returns a non-nil value, the panic is stopped and execution continues after the deferred call.
- **Use case**: to make a library safer, or to convert a panic into an error before returning.

```go
func safer() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("caught panic:", r)
        }
    }()
    panic("oh no")
}
```

> **Key point:** `panic` is for raising an error condition; `recover` lets you handle it if invoked in a deferred call.

---

## Value Receiver vs. Pointer Receiver

When you define methods on a type:

### **Value receiver**

- Method operates on a copy of the value.
- Changes made inside the method **do not** affect the original.
- Suitable when the type is small or immutable, e.g.:

```go
func (p Person) String() string { // receiver is a copy
    return p.Name
}
```

- The method can be called on both value and pointer instances.

### **Pointer receiver**

- Receiver is a pointer (`*Type`), so the method works with the original value.
- Modifications inside the method **affect** the caller's instance.
- Efficient for large structs (avoids copying) and when you need to modify the receiver:

```go
func (p *Person) SetName(n string) {
    p.Name = n
}
```

- Can only be called with pointers (though Go will auto-take the address when calling with a value).

> **Guidelines:**
>
> - Use pointer receivers when the method needs to mutate state or if the type is large.
> - Keep receiver type consistent across all methods of a given type.
