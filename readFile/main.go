package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer to hold chunks of data (e.g., 100 bytes)
	buffer := make([]byte, 100)

	for {
		// Read a chunk into the buffer
		n, err := file.Read(buffer)

		// Process the bytes read (from index 0 to n-1)
		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}

		// Check if the end of the file has been reached
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
	}
}
