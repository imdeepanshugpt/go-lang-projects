# Runes vs Bytes in Go

Go provides two distinct aliases for integer types that are commonly used when
working with text: `byte` and `rune`. Both are just type aliases, but they
carry important semantic meaning.

## Definitions

- **`byte`**
  - Alias for `uint8`.
  - Represents a single 8‑bit unit.
  - Commonly used for raw binary data and UTF-8 code units.

- **`rune`**
  - Alias for `int32`.
  - Represents a Unicode code point (a character).
  - Decoded from UTF-8 and can hold any valid Unicode value.

## Why not just cast?

Because `byte` and `rune` are aliases, you can always convert between them and
other integer types:

```go
var b byte = 'A'         // rune literal converted to byte
var r rune = rune(b)     // explicit conversion
```

However, the purpose of the separate names is to express _intent_ and to
match the behavior of Go’s string and text APIs:

1. **Documentation by type** – a function taking a `rune` signals it expects a
   character, whereas `byte` means a raw code unit or binary value.
2. **Correct semantics** – iterating over a string with `range` yields runes;
   indexing yields bytes. Using the appropriate alias reduces confusion.
3. **Avoid bugs** – handling multi-byte UTF-8 characters as a sequence of
   `uint8` values leads to incorrect counts, slicing, and processing logic.
4. **Standard library consistency** – many packages expose APIs using `rune`
   (e.g. `strings.IndexRune`) or `byte` (e.g. `bytes.Buffer`), so matching
   those types helps interoperability.

## Examples

### Bytes from a string

```go
s := "café"
for i := 0; i < len(s); i++ {
    fmt.Printf("%d: %x\n", i, s[i])
}
// prints the raw UTF-8 bytes: 63 61 66 c3 a9
```

### Runes via range

```go
for i, r := range "café" {
    fmt.Printf("%d: %c (U+%04X)\n", i, r, r)
}
// 0: c (U+0063)
// 1: a (U+0061)
// 2: f (U+0066)
// 3: é (U+00E9)
```

### Iteration differences

- `range` decodes UTF-8 and returns runes.
- `len(s)` and indexing operate on bytes.

## Summary

| Name | Underlying Type | Size    | Meaning                   |
| ---- | --------------- | ------- | ------------------------- |
| byte | `uint8`         | 1 byte  | raw data / UTF-8 unit     |
| rune | `int32`         | 4 bytes | Unicode code point (char) |

Use `byte` when working with binary data or when you need precise control
over UTF-8 code units. Use `rune` when processing text in terms of
characters.

Although you can cast between them, choosing the right alias makes your code
clearer and reduces the chance of Unicode-related bugs.
