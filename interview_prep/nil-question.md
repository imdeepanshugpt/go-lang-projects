# Go Nil Behavior — Complete Guide

A comprehensive reference for understanding `nil` in Go, covering all the common traps, edge cases, and interview questions.

---

## Table of Contents

1. [Types That Can Be Nil](#types-that-can-be-nil)
2. [Nil Pointer Dereference](#1-nil-pointer-dereference)
3. [Nil Slice vs Empty Slice](#2-nil-slice-vs-empty-slice)
4. [Nil Map Write Panic](#3-nil-map-write-panic)
5. [Nil Channel Deadlock](#4-nil-channel-deadlock)
6. [Interface Nil Trap](#5-interface-nil-trap-️-famous-interview-question)
7. [Nil Function Call](#6-nil-function-call)
8. [Nil Slice Append](#7-nil-slice-append-works)
9. [Nil Interface Proper Check](#8-nil-interface-proper-check)
10. [Nil Struct Pointer Method Call](#9-nil-struct-pointer-method-call)
11. [Nil vs Zero Value Struct](#10-nil-vs-zero-value-struct)

---

## Types That Can Be Nil

| Type       | Can Be Nil? |
| ---------- | ----------- |
| Pointer    | ✅ Yes      |
| Slice      | ✅ Yes      |
| Map        | ✅ Yes      |
| Channel    | ✅ Yes      |
| Interface  | ✅ Yes      |
| Function   | ✅ Yes      |
| Struct     | ❌ No       |
| Primitives | ❌ No       |

> **Rule of thumb:** Only reference types and interfaces can be nil. Value types (structs, int, bool, etc.) cannot.

---

## 1. Nil Pointer Dereference

**Code:**

```go
package main

import "fmt"

func main() {
    var p *int
    fmt.Println(*p) // panic!
}
```

**Output:**

```
panic: runtime error: invalid memory address or nil pointer dereference
```

**Why?**
`p` is declared but never assigned — it defaults to `nil`. Dereferencing `*p` tries to read memory at address `0x0`, which the OS protects. The runtime catches this and panics.

**Fix:**

```go
if p != nil {
    fmt.Println(*p)
}
```

---

## 2. Nil Slice vs Empty Slice

**Code:**

```go
var s []int           // nil slice
fmt.Println(s == nil) // true
fmt.Println(len(s))   // 0

s2 := []int{}           // empty slice
fmt.Println(s2 == nil)  // false
fmt.Println(len(s2))    // 0
```

**Output:**

```
true
0
false
0
```

**Why?**
Internally, a slice is a struct with three fields: `(pointer, length, capacity)`.

| Slice          | Pointer | Length | Capacity |
| -------------- | ------- | ------ | -------- |
| `var s []int`  | `nil`   | `0`    | `0`      |
| `s := []int{}` | non-nil | `0`    | `0`      |

A nil slice and an empty slice behave identically for most operations (`len`, `cap`, `range`, `append`), but they differ in nil equality. Prefer `len(s) == 0` over `s == nil` in most practical checks.

---

## 3. Nil Map Write Panic

**Code:**

```go
var m map[string]int
m["age"] = 30 // panic!
```

**Output:**

```
panic: assignment to entry in nil map
```

**Why?**
A nil map has no underlying hash table allocated. Reading from a nil map is safe (returns the zero value), but writing requires an initialized map.

**Fix:**

```go
m := make(map[string]int)
m["age"] = 30
```

> **Note:** Reading from a nil map is safe: `fmt.Println(m["age"])` → `0`, no panic.

---

## 4. Nil Channel Deadlock

**Code:**

```go
var ch chan int
ch <- 5 // blocks forever
```

**Output:**

```
fatal error: all goroutines are asleep - deadlock!
```

**Why?**
A nil channel is never ready for communication — sends and receives on it block forever. Since no other goroutine can unblock it, the runtime detects a deadlock.

**Fix:**

```go
ch := make(chan int, 1) // buffered, or use unbuffered with a goroutine
ch <- 5
```

> **Useful pattern:** Nil channels are intentionally used to disable a `select` case dynamically.

---

## 5. Interface Nil Trap ⚠️ Famous Interview Question

**Code:**

```go
package main

import "fmt"

type User struct{}

func main() {
    var u *User = nil
    var i interface{} = u

    fmt.Println(i == nil) // false!
}
```

**Output:**

```
false
```

**Why?**
An interface value internally holds two fields:

```
interface{ type, value }
```

When you assign `u` (a `*User` nil pointer) to `i`, the interface becomes:

```
{ type: *User, value: nil }
```

The interface itself is **not** nil because it has a concrete type. Only an interface with **both** type and value set to nil is truly nil:

```go
var i interface{} = nil  // { type: nil, value: nil } → truly nil
```

This is one of the most famous Go gotchas — it catches even experienced developers.

---

## 6. Nil Function Call

**Code:**

```go
var f func()
f() // panic!
```

**Output:**

```
panic: runtime error: invalid memory address or nil pointer dereference
```

**Why?**
Function variables are pointers under the hood. A `var f func()` declaration sets `f` to nil. Calling a nil function pointer causes a panic.

**Fix:**

```go
if f != nil {
    f()
}
```

---

## 7. Nil Slice Append (Works!)

**Code:**

```go
var s []int
s = append(s, 10)
fmt.Println(s) // [10]
```

**Output:**

```
[10]
```

**Why?**
`append` handles nil slices gracefully — it allocates a new underlying array when needed. This is by design in Go, making it safe to declare a slice and immediately start appending without initialization.

> **Contrast with maps:** You must `make` a map before writing, but nil slices work fine with `append`.

---

## 8. Nil Interface Proper Check

**Problem:**

```go
func isNil(i interface{}) bool {
    return i == nil // fails when a nil pointer is wrapped in interface
}
```

This returns `false` when called with a nil pointer (see the [Interface Nil Trap](#5-interface-nil-trap-️-famous-interview-question) above).

**Basic fix using reflection:**

```go
import "reflect"

func isNil(i interface{}) bool {
    if i == nil {
        return true
    }
    v := reflect.ValueOf(i)
    return v.Kind() == reflect.Ptr && v.IsNil()
}
```

**Complete fix (handles all nilable types):**

```go
import "reflect"

func isNil(i interface{}) bool {
    if i == nil {
        return true
    }
    v := reflect.ValueOf(i)
    switch v.Kind() {
    case reflect.Ptr, reflect.Slice, reflect.Map,
         reflect.Chan, reflect.Func, reflect.Interface:
        return v.IsNil()
    }
    return false
}
```

**Why the complete version is better:**
The basic version only handles pointer types. The complete version covers all nilable kinds: pointers, slices, maps, channels, functions, and interfaces.

---

## 9. Nil Struct Pointer Method Call

**Code:**

```go
type User struct {
    name string
}

func (u *User) Print() {
    fmt.Println("hello")
}

func main() {
    var u *User
    u.Print() // works!
}
```

**Output:**

```
hello
```

**Why?**
In Go, calling a method on a nil pointer is legal — the nil pointer is simply passed as the receiver. As long as the method body doesn't dereference the pointer, it works fine.

**When it panics:**

```go
func (u *User) PrintName() {
    fmt.Println(u.name) // panic: nil pointer dereference
}
```

Accessing a field (`u.name`) dereferences the pointer, causing a panic.

**Practical pattern:** You can use nil receiver methods to provide safe defaults:

```go
func (u *User) Name() string {
    if u == nil {
        return "anonymous"
    }
    return u.name
}
```

---

## 10. Nil vs Zero Value Struct

**Code:**

```go
var u User
fmt.Println(u == nil) // compilation error!
```

**Output:**

```
invalid operation: u == nil (mismatched types User and untyped nil)
```

**Why?**
Structs are value types in Go. They are never nil — when you declare `var u User`, Go zero-initializes all fields. Comparing a value type to `nil` is a type error caught at compile time.

**Fix — use a pointer:**

```go
var u *User
fmt.Println(u == nil) // true ✅
```

---

## Quick Reference Cheat Sheet

| Scenario                            | Result              | Fix                          |
| ----------------------------------- | ------------------- | ---------------------------- |
| Dereference nil pointer             | panic               | Check `!= nil` first         |
| Read from nil map                   | zero value, safe    | —                            |
| Write to nil map                    | panic               | `make(map[K]V)`              |
| Send on nil channel                 | blocks forever      | `make(chan T)`               |
| Nil pointer in interface `== nil`   | `false`             | Use `reflect.ValueOf`        |
| Call nil function                   | panic               | Check `!= nil` first         |
| `append` to nil slice               | works               | —                            |
| Call method on nil pointer receiver | works (if no deref) | Guard with nil check in body |
| Compare struct to nil               | compile error       | Use pointer `*Struct`        |

---

## Summary

Go's `nil` is not a universal "no value" like in some other languages — its behavior depends entirely on the type. The golden rules:

1. **Always initialize maps and channels** with `make`.
2. **Check pointers and functions** before use with `!= nil`.
3. **Never trust `i == nil`** when an interface might hold a nil pointer — use reflection.
4. **Nil slices are safe to use** — `append`, `len`, `range` all work.
5. **Structs are never nil** — only pointers to structs can be.
