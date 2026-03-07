package main

import "fmt"

// makeNew demonstrates the difference between make and new by returning
// examples of each usage.
func makeNew() {
	// new: allocates zeroed storage for a type and returns a pointer.
	// Works with any type, but the memory is not otherwise initialized.
	pInt := new(int) // *int, points to 0
	*pInt = 42       // we can set the value via the pointer

	pSlice := new([]int)     // *[]int, points to a nil slice
	*pSlice = []int{1, 2, 3} // we must create the slice separately

	fmt.Printf("using new: pInt=%p val=%d, pSlice=%p val=%v\n", pInt, *pInt, pSlice, *pSlice)

	// make: used only for slices, maps, and channels.
	// It allocates the underlying data structure and returns it (not a pointer).
	s := make([]int, 3, 5)    // len 3, cap 5, initialized with zeroes
	m := make(map[string]int) // empty map ready for use
	c := make(chan string, 2) // buffered channel of capacity 2

	s[0] = 100
	m["foo"] = 1
	c <- "hello"

	fmt.Printf("using make: s=%v len=%d cap=%d, m=%v, c=<buffered>\n",
		s, len(s), cap(s), m)
}

func main() {
	makeNew()
}

// Question: What is the difference between make and new in Go?
// In Go, `make` and `new` are built-in functions that serve different
// purposes when it comes to memory allocation and initialization.
//
// 1. `new`: The `new` function is used to allocate memory for a variable of a
//    specific type and returns a pointer to that memory. It initializes the
//    allocated memory to the zero value of the type. For example, `new(int)`
//    will allocate memory for an integer and return a pointer to it,
//    initialized to 0.
//
// 2. `make`: The `make` function is used to create and initialize slices,
//    maps, and channels. It allocates and initializes the underlying data
//    structure for these types. For example, `make([]int, 5)` will create a
//    slice of integers with a length of 5 and a capacity of 5, initialized
//    with zero values.
//
// In summary, `new` is used for allocating memory for basic types and
// returns a pointer, while `make` is used for creating and initializing
// slices, maps, and channels, returning the initialized value directly.
