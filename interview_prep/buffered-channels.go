package main

import (
	"fmt"
	"time"
)

// Example 1: Basic Buffered Channel
// A buffered channel has capacity; sender blocks only when buffer is full
func example1_BasicBufferedChannel() {
	fmt.Println("=== Example 1: Basic Buffered Channel ===")

	// Create a buffered channel with capacity 2
	ch := make(chan int, 2)

	// Send without goroutine (because buffer has space)
	ch <- 1
	ch <- 2
	fmt.Println("Sent 2 values to buffer")

	// This would block because buffer is full:
	// ch <- 3 // blocks here

	// Receive values
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Println()
}

// Example 2: Buffer Capacity vs Unbuffered
func example2_BufferedVsUnbuffered() {
	fmt.Println("=== Example 2: Buffered vs Unbuffered ===")

	// Unbuffered - sender blocks immediately until receiver is ready
	fmt.Println("Unbuffered Channel:")
	ch1 := make(chan int)
	go func() {
		fmt.Println("  Goroutine: sending value")
		ch1 <- 42
		fmt.Println("  Goroutine: value sent")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Main: receiving")
	fmt.Printf("  Main: received %d\n", <-ch1)

	// Buffered - sender doesn't block if buffer has space
	fmt.Println("\nBuffered Channel (capacity 2):")
	ch2 := make(chan int, 2)
	go func() {
		fmt.Println("  Goroutine: sending value 1")
		ch2 <- 1
		fmt.Println("  Goroutine: value 1 sent (no wait!)")
		fmt.Println("  Goroutine: sending value 2")
		ch2 <- 2
		fmt.Println("  Goroutine: value 2 sent (no wait!)")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Main: receiving")
	fmt.Printf("  Main: received %d\n", <-ch2)
	fmt.Printf("  Main: received %d\n", <-ch2)
	fmt.Println()
}

// Example 3: Buffered Channel as Queue
func example3_BufferedAsQueue() {
	fmt.Println("=== Example 3: Buffered Channel as Queue ===")

	// Create a queue with capacity 5
	queue := make(chan string, 5)

	// Enqueue operations
	queue <- "task1"
	queue <- "task2"
	queue <- "task3"

	fmt.Printf("Queue length: %d, capacity: %d\n", len(queue), cap(queue))

	// Dequeue operations
	fmt.Printf("Dequeued: %s\n", <-queue)
	fmt.Printf("Dequeued: %s\n", <-queue)

	fmt.Printf("Queue length after dequeue: %d\n\n", len(queue))
}

// Example 4: Checking Buffer Availability
func example4_CheckingBuffer() {
	fmt.Println("=== Example 4: Checking Buffer Length ===")

	ch := make(chan int, 3)

	// Fill buffer
	ch <- 10
	ch <- 20

	fmt.Printf("Values in buffer: %d\n", len(ch))
	fmt.Printf("Buffer capacity: %d\n", cap(ch))
	fmt.Printf("Empty slots: %d\n\n", cap(ch)-len(ch))
}

// Example 5: Fan-Out Pattern with Buffered Channel
func example5_FanOut() {
	fmt.Println("=== Example 5: Fan-Out Pattern ===")

	// Main task channel (buffered)
	tasks := make(chan int, 5)

	// Send multiple tasks at once (no blocking due to buffer)
	for i := 1; i <= 5; i++ {
		tasks <- i
	}

	// Process tasks
	for i := 0; i < 5; i++ {
		task := <-tasks
		fmt.Printf("Processing task: %d\n", task)
	}
	fmt.Println()
}

// Example 6: Buffered Channel with Timeout
func example6_BufferedWithTimeout() {
	fmt.Println("=== Example 6: Buffered with Timeout ===")

	results := make(chan string, 10)

	// Simulate slow producer
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(500 * time.Millisecond)
			results <- fmt.Sprintf("Result %d", i)
		}
	}()

	// Consume with timeout
	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout waiting for results")
			return
		}
	}
}

// Example 7: Semaphore Pattern (Buffered Channel)
// Limit concurrency using buffered channel
func example7_Semaphore() {
	fmt.Println("=== Example 7: Semaphore Pattern ===")

	// Semaphore: allow max 3 concurrent goroutines
	semaphore := make(chan struct{}, 3)

	for i := 1; i <= 7; i++ {
		go func(id int) {
			semaphore <- struct{}{} // Acquire
			fmt.Printf("Goroutine %d acquired slot\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d releasing slot\n", id)
			<-semaphore // Release
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println()
}

// Example 8: When to use Buffered Channels
func example8_WhenToUseBuffered() {
	fmt.Println("=== Example 8: When to Use Buffered Channels ===")

	// Use case 1: Decoupling producer and consumer speeds
	fmt.Println("Use case 1: Speed Mismatch")
	slow := make(chan int, 5) // Buffer absorbs fast producer

	go func() {
		for i := 1; i <= 5; i++ {
			slow <- i
			fmt.Printf("Produced: %d\n", i)
		}
	}()

	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 5; i++ {
		val := <-slow
		fmt.Printf("Consumed: %d\n", val)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
}

// Example 9: Multiple Senders to Buffered Channel
func example9_MultipleSenders() {
	fmt.Println("=== Example 9: Multiple Senders ===")

	results := make(chan string, 10)

	// Multiple goroutines sending to same buffered channel
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 2; j++ {
				msg := fmt.Sprintf("Sender %d, message %d", id, j)
				results <- msg
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(500 * time.Millisecond)

	// Collect all results
	for i := 0; i < 6; i++ {
		fmt.Println(<-results)
	}
	fmt.Println()
}

// Example 10: Draining Buffered Channel
func example10_DrainChannel() {
	fmt.Println("=== Example 10: Draining Buffered Channel ===")

	ch := make(chan int, 5)

	// Fill channel
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	fmt.Printf("Channel has %d items\n", len(ch))

	// Drain all values
	for len(ch) > 0 {
		val := <-ch
		fmt.Printf("Drained: %d\n", val)
	}

	fmt.Printf("Channel drained, remaining: %d\n\n", len(ch))
}

func main() {
	example1_BasicBufferedChannel()
	example2_BufferedVsUnbuffered()
	example3_BufferedAsQueue()
	example4_CheckingBuffer()
	example5_FanOut()
	example6_BufferedWithTimeout()
	example7_Semaphore()
	example8_WhenToUseBuffered()
	example9_MultipleSenders()
	example10_DrainChannel()
}
