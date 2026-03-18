package main

import "fmt"

/*
- Given a large integer as an integer array where digit[i] is the digit of the integer.
- The digits are ordered from most signifigant to least signifigant in left to right order.
- Increment the large integer by one and return the resulting array of digits.
*/

func main() {
	fmt.Println("-----PLUS ONE-----")

	digits := []int{1, 2, 3, 4}
	result := plusOne(digits)

	fmt.Println("Result:", result)
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
