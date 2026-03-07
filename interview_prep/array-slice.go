package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // This creates a slice from index 1 to 3 (4 is exclusive)
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	// Modifying the slice will affect the original array
	slice[0] = 10
	fmt.Println("Modified Slice:", slice)
	fmt.Println("Modified Array:", arr)
}

/*
What is the difference between array and slice in Go?
In Go, an array is a fixed-size collection of elements of the same type. The size of an array is determined at compile time and cannot be changed. An array is defined with a specific length, and its elements are stored contiguously in memory.

On the other hand, a slice is a dynamically-sized, flexible view into the elements of an array. A slice does not store any data itself; it just describes a section of an underlying array. A slice has three components: a pointer to the underlying array, the length of the slice (the number of elements in the slice), and the capacity (the maximum number of elements the slice can grow to).

In summary, the main differences between arrays and slices in Go are:
1. Size: Arrays have a fixed size, while slices can grow and shrink dynamically.
2. Memory: Arrays store their elements directly, while slices reference an underlying array.
3. Usage: Slices are more commonly used in Go due to their flexibility and ease of use compared to arrays.

How does slice capacity work?
The capacity of a slice in Go refers to the total number of elements that the slice can hold before it needs to be resized. When you create a slice, it has an initial length (the number of elements currently in the slice) and a capacity (the maximum number of elements it can hold without resizing). When you append elements to a slice and it exceeds its current capacity, Go automatically allocates a new underlying array with a larger capacity, copies the existing elements to the new array, and updates the slice to reference the new array. The new capacity is typically doubled to minimize future reallocations.

What happens when you append to a slice?
When you append to a slice in Go, the following steps occur:
1. If the length of the slice is less than its capacity, the new element is added to the existing underlying array, and the length of the slice is increased by one.
2. If the length of the slice equals its capacity, Go allocates a new underlying array with a larger capacity (usually doubling the current capacity), copies the existing elements from the old array to the new array, and then adds the new element to the new array. The slice is updated to reference the new array, and its length and capacity are updated accordingly.

This process allows slices to grow dynamically while managing memory efficiently.
*/
