package main

import (
	"fmt"
	"time"
)

// Example 1: Basic Unbuffered Channel
// An unbuffered channel blocks until both sender and receiver are ready
func example1_BasicChannel() {
	fmt.Println("=== Example 1: Basic Unbuffered Channel ===")

	// Create an unbuffered channel of type int
	ch := make(chan int)

	// Goroutine to send data
	go func() {
		fmt.Println("Sending 42 to channel...")
		ch <- 42 // Send blocks until someone receives
		fmt.Println("Value sent!")
	}()

	// Receive data
	fmt.Println("Waiting to receive...")
	value := <-ch // Receive blocks until someone sends
	fmt.Printf("Received: %d\n\n", value)
}

// Example 2: Ping-Pong with Channels
func example2_PingPong() {
	fmt.Println("=== Example 2: Ping-Pong Communication ===")

	ping := make(chan string)
	pong := make(chan string)

	go func() {
		ping <- "ping"
	}()

	go func() {
		msg := <-ping
		fmt.Printf("Received: %s\n", msg)
		pong <- "pong"
	}()

	msg := <-pong
	fmt.Printf("Received: %s\n\n", msg)
}

// Example 3: Worker Pattern
func example3_WorkerPattern() {
	fmt.Println("=== Example 3: Worker Pattern ===")

	jobs := make(chan int)
	results := make(chan string)

	// Worker goroutine
	go func() {
		for job := range jobs { // Blocks until channel is closed
			fmt.Printf("Processing job: %d\n", job)
			time.Sleep(1 * time.Millisecond)
			results <- fmt.Sprintf("Result of %d", job)
		}
	}()

	// Send jobs
	go func() {
		for i := 1; i <= 3; i++ {
			jobs <- i
		}
		close(jobs) // Signal done sending
	}()

	// Receive results
	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
	fmt.Println()
}

// Example 4: Multiple Goroutines with WaitGroup Pattern
func example4_MultipleWorkers() {
	fmt.Println("=== Example 4: Multiple Workers ===")

	tasks := make(chan int)
	results := make(chan string)

	// Start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go func(workerID int) {
			for task := range tasks {
				fmt.Printf("Worker %d processing task %d\n", workerID, task)
				results <- fmt.Sprintf("Worker %d completed task %d", workerID, task)
			}
		}(w)
	}

	// Send tasks
	go func() {
		for i := 1; i <= 5; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// Receive results
	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
	fmt.Println()
}

// Example 5: Select with Channels (Multiplexing)
func example5_Select() {
	fmt.Println("=== Example 5: Select - Multiplexing ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "two"
	}()

	// Wait for first available message
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}
	fmt.Println()
}

// Example 6: Timeout with Select
func example6_Timeout() {
	fmt.Println("=== Example 6: Timeout with Select ===")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "delayed message"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! No message received within 1 second")
	}
	fmt.Println()
}

// Example 7: Closing Channels
func example7_CloseChannel() {
	fmt.Println("=== Example 7: Closing Channels ===")

	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // Signal that no more values will be sent
	}()

	// Iterate until channel is closed
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}

	// After closure, receive returns zero value
	val, ok := <-ch
	fmt.Printf("After close - Value: %d, OK: %v\n\n", val, ok)
}

func main() {
	example1_BasicChannel()
	example2_PingPong()
	example3_WorkerPattern()
	example4_MultipleWorkers()
	example5_Select()
	example6_Timeout()
	example7_CloseChannel()
}
