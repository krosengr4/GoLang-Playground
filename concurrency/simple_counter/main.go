package main

import "fmt"

// Create a program with a goroutine that counts from
// 1 to 10, sending each number through a channel to
// the main goroutine, which prints them.

func main() {
	// Create a channel to receive integers
	numbers := make(chan int)

	// Launch a goroutine that counts from 1 to 10
	go func() {
		for i := 1; i <= 10; i++ {
			numbers <- i // Send number to channel
		}
		close(numbers) // Close the channel when done
	}()

	// Receive and print numbers from the channel
	for num := range numbers {
		fmt.Println("Received:", num)
	}

	fmt.Println("All numbers received!")
}
