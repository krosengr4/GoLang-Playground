package main

import "fmt"

// Good ol fizzbuzz
// Print the integers 1 - 100
// for multiples of 3, print fizz instead of the number
// for multiples of 5, print buzz instead of the number
// And for multiples of 3 and 5... you guessed it, print fizzbuzz instead of the number

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("fizzbuzz!")
		case i%5 == 0:
			fmt.Println("buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}
	}
}
