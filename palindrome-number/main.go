package main

import (
	"strings"
	"fmt"
)

/*
- Given an integer, return true if it is a palindrome, and false otherwise
- An integer is a palindrome when it reads the same forward and backward.
*/

func main() {
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("\t------PALINDROME NUMBER-----")
	fmt.Println(strings.Repeat("-", 55))

	result := isPalindrome(1441)
	fmt.Println("Result: ", result)
}

func isPalindrome(x int) bool {
	// Edge cases
	if x < 0 || (x % 10 == 0 && x != 0){
		return false
	}

	reverseNum := 0

	// Reverse only the back half of the number
	for x > reverseNum {
		// Modulo of 10 will always give last digit
		reverseNum = (reverseNum * 10) + (x % 10)

		// remove last digit from x by dividing by 10
		x /= 10
	}

	fmt.Println(reverseNum)

	return x == reverseNum || x == reverseNum/10
}
