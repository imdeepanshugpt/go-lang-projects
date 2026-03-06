package main

import (
	"fmt"
)

func main() {
	greeting := "Hi There!"

	go (func() {
		fmt.Println(greeting)
	})()
}
