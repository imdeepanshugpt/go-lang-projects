# Arrays vs Slices in Go

This document provides a quick reference for the syntactic differences between arrays and slices in Go, how slice capacity grows, and a small program that illustrates backing-array behavior.

---

## 1. Creating an **array**

```go
// an array of five ints – length is part of the type
var a [5]int                   // zero‑valued: [0 0 0 0 0]
b := [3]string{"x", "y", "z"}  // literal, length 3

// the size is fixed at compile time and cannot be changed:
a[0] = 42
// a = append(a, 1)    // ✗ compile error: cannot append to array
```

- Size is part of the type (`[5]int` ≠ `[6]int`)
- Elements are stored **directly** in the variable; the variable _is_ the array.
- Copying an array copies all elements.

---

## 2. Creating a **slice**

```go
// slice literal – no size in the type
s := []int{1, 2, 3}        // len 3, cap 3

// from an existing array/ slice
arr := [5]int{10, 20, 30, 40, 50}
s2 := arr[1:4]             // len 3, cap 4 (from index 1 to end of arr)
```

- A slice is a descriptor with a pointer, length and capacity:

  ```
  type sliceHeader struct {
      ptr uintptr
      len int
      cap int
  }
  ```

- It **references** an underlying array; it doesn’t own its storage.
- Slices are growable; arrays are not.

---

## 3. How capacity grows

Appending to a slice behaves as follows:

1. If `len < cap`, the new element is stored in the same backing array.
2. If `len == cap`, Go allocates a new array (usually double the old capacity, with some heuristics),
   copies the old elements, and returns a slice pointing at the new array.

The growth policy is implementation‑defined but typically:

```text
cap → cap*2    (for small slices)
cap → cap + n  (when cap is already large)
```

---

## 4. Example program showing addresses of the backing array

```go
package main

import "fmt"

func main() {
    s := make([]int, 0, 1)            // len=0, cap=1
    fmt.Printf("initial: len=%d cap=%d ptr=%p\n", len(s), cap(s), s)

    for i := 0; i < 5; i++ {
        s = append(s, i)
        fmt.Printf("after append %d: len=%d cap=%d ptr=%p -> %v\n",
            i, len(s), cap(s), s, s)
    }
}
```

Sample output:

```
initial: len=0 cap=1 ptr=0xc0000140a0
after append 0: len=1 cap=1 ptr=0xc0000140a0 -> [0]
after append 1: len=2 cap=2 ptr=0xc0000140b0 -> [0 1]
after append 2: len=3 cap=4 ptr=0xc0000140d0 -> [0 1 2]
after append 3: len=4 cap=4 ptr=0xc0000140d0 -> [0 1 2 3]
after append 4: len=5 cap=8 ptr=0xc0000140e0 -> [0 1 2 3 4]
```

Notice the `ptr` value changes when capacity is exceeded – that’s the new underlying array.

> ⚠️ The addresses are printed for demonstration; they will differ each run.

---

## 5. Memory layout (conceptual)

```
array value: [4]int{…}

slice value:
  header { ptr --------v
           len cap }
                |
                v
         [0 1 2 3 4 5 … ]  // underlying array
```

- The slice header is a small struct on the stack (or wherever the slice variable lives).
- The **data** lives in a separate heap‑allocated array when the slice is grown.

---

Slices are the idiomatic, flexible way to work with sequences in Go; arrays are useful for fixed‑size buffers and as underlying storage for slices.
