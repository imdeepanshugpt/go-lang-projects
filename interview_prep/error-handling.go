package main

import (
	"errors"
	"fmt"
)

// f returns a simple error created with the errors package.
func f() error {
	return errors.New("something went wrong")
}

// g demonstrates using fmt.Errorf to wrap an existing error.
func g(path string) error {
	// pretend we tried to open a file and got an error
	var err error = errors.New("permission denied")
	if err != nil {
		// wrap the underlying error with context
		return fmt.Errorf("failed to open %s: %w", path, err)
	}
	return nil
}

func main() {
	// Example of using f
	if err := f(); err != nil {
		fmt.Println("Error from f:", err)
	}

	// Example of using g
	if err := g("/path/to/file"); err != nil {
		fmt.Println("Error from g:", err)
	}
}

/*
What is error handling in Go and how does it differ from other languages?
In Go, error handling is done using the built-in `error` type, which is an interface that represents an error condition. Functions that can fail typically return an `error` as the last return value, allowing the caller to check for errors explicitly. This approach encourages developers to handle errors immediately and clearly, rather than relying on exceptions or other mechanisms.

In contrast to languages that use exceptions (like Java or Python), Go does not have a try-catch mechanism. Instead, it promotes a more explicit and straightforward way of handling errors, where the programmer must check for errors at each step. This leads to clearer code and reduces the chances of unhandled exceptions, as errors are treated as regular values that can be passed around and handled as needed.

Additionally, Go provides the `fmt.Errorf` function, which allows developers to wrap existing errors with additional context, making it easier to trace the source of an error when it occurs. This is a powerful feature that helps in debugging and understanding the flow of errors in a Go program.
How do you create and handle custom errors in Go?
In Go, you can create custom errors by defining a new type that implements the `error` interface. The `error` interface requires a single method, `Error() string`, which returns a string representation of the error. Here's an example of how to create and handle custom errors in Go:
```go
package main

import (
	"fmt"
)

// MyError is a custom error type that includes additional context.
type MyError struct {
	Code    int
	Message string
}

// Error implements the error interface for MyError.
func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// doSomething simulates a function that can return a custom error.
func doSomething() error {
	return &MyError{Code: 404, Message: "Resource not found"}
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("Custom error occurred:", err)
	}
}
In this example, we define a custom error type `MyError` that includes a code and a message. The `Error()` method formats these fields into a string. The `doSomething` function simulates an operation that returns a `MyError`. In the `main` function, we call `doSomething` and check for errors, printing the custom error message if it occurs. This approach allows you to create rich error types that can carry additional information beyond just a string message.
How does error wrapping work in Go and how can it be used to provide more context about an error?
Error wrapping in Go is a technique that allows you to add additional context to an existing error while preserving the original error information. This is typically done using the `fmt.Errorf` function with the `%w` verb, which indicates that the error should be wrapped.

When you wrap an error, you can provide a new message that gives more context about where or why the error occurred, while still retaining access to the original error. This is useful for debugging and understanding the flow of errors in your application.

Here's an example of how error wrapping works in Go:
```gopackage main

import (
	"errors"
	"fmt"
)

// g demonstrates using fmt.Errorf to wrap an existing error.
func g(path string) error {
	// pretend we tried to open a file and got an error
	var err error = errors.New("permission denied")
	if err != nil {
		// wrap the underlying error with context
		return fmt.Errorf("failed to open %s: %w", path, err)
	}
	return nil
}

func main() {
	if err := g("/path/to/file"); err != nil {
		fmt.Println("Error from g:", err)

		// Unwrap the error to get the original error
		if errors.Is(err, errors.New("permission denied")) {
			fmt.Println("The original error is permission denied")
		}
	}
}
In this example, the `g` function simulates an operation that fails with a "permission denied" error. We use `fmt.Errorf` to wrap this error with additional context about the file path that caused the error. In the `main` function, we check for the error and print it. We can also use `errors.Is` to check if the original error is "permission denied", demonstrating how we can access the underlying error information even after wrapping it with additional context.

*/
