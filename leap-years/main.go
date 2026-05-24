package main

import "fmt"

/*
- In this kata you should simply determine, whether a given year is a leap year or not.
- A leap year is:
- Years divisible by 4 are leap years
- Not divisible by 100 UNLESS divisible by 400
*/

func main() {
	year := 2020
	result := isLeapYear(year)
	fmt.Println(result)

}

func isLeapYear(year int) bool {
	divisibleByFour := year%4 == 0
	isCentury := year%100 == 0
	isDivisibleByFourHundred := year%400 == 0

	if divisibleByFour && (!isCentury || isDivisibleByFourHundred) {
		return true
	} else {
		return false
	}

}
