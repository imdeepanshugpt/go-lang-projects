package main

import "fmt"

func main() {
	var p *int
	fmt.Println(*p)
}

// panic: runtime error: invalid memory address or nil pointer dereference
// Why?

// p is nil, and dereferencing *p tries to access memory that doesn't exist.

// Correct check:

// if p != nil {
// 	fmt.Println(*p)
// }

// Nil Slice vs Empty Slice
// Question
// var s []int

// fmt.Println(s == nil)
// fmt.Println(len(s))
// Output
// true
// 0

// Explanation:

// A nil slice behaves like an empty slice, but internally they differ.

/*
Example:
s := []int{}
Output:
s == nil -> false
len(s) -> 0

Nil Map Write Panic
Question
var m map[string]int
m["age"] = 30
Result
panic: assignment to entry in nil map
Why?

Maps must be initialized.

Correct:

m := make(map[string]int)
m["age"] = 30

Nil Channel Deadlock
Question
var ch chan int
ch <- 5
Result
deadlock

Because nil channels block forever.

Correct:

ch := make(chan int)

Interface Nil Trap (Very Famous Question):
Question
package main
import "fmt"

type User struct{}

func main() {

	var u *User = nil
	var i interface{} = u

	fmt.Println(i == nil)
}
Output
false
Why?

Interface internally stores:

(type, value)

So here:

type = *User
value = nil

Interface itself is not nil.

This is one of the most famous Go interview traps.

Nil Function Call
Question
var f func()
f()
Result
panic: runtime error

Because function pointer is nil.

Correct check:

if f != nil {
	f()
}

Nil Slice Append (Works!)
Question
var s []int
s = append(s, 10)

fmt.Println(s)
Output
[10]

Unlike maps, nil slices can grow dynamically.

Nil Interface Proper Check
Problem
func isNil(i interface{}) bool {
	return i == nil
}

This fails when pointer is inside interface.

Proper check requires reflection.

Example:

import "reflect"

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)

	return v.Kind() == reflect.Ptr && v.IsNil()
}

Nil Struct Pointer Method Call
Question
type User struct{}

func (u *User) Print() {
	fmt.Println("hello")
}

func main() {
	var u *User
	u.Print()
}
Output
hello

Why?

Method does not dereference pointer, so it's allowed.

But if code used:

fmt.Println(u.name)

Nil Struct Pointer Method Call
Question
type User struct{}

func (u *User) Print() {
	fmt.Println("hello")
}

func main() {
	var u *User
	u.Print()
}
Output
hello

Why?

Method does not dereference pointer, so it's allowed.

But if code used:

fmt.Println(u.name)

It would panic.

Nil vs Zero Value Struct
Question
var u User

fmt.Println(u == nil)
Result

Compilation error:

invalid operation: u == nil

Why?

Structs cannot be nil.

Only pointers can be nil.

Correct:

var u *User

What types can be nil in Go?

Answer:

Pointer
Slice
Map
Channel
Interface
Function

Structs and primitive types cannot be nil.
