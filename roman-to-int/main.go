package main

import (
	"fmt"
	"strings"
)

/*
- Given a Roman Numeral (ex: XXVII) convert it to and integer (ex: 17)
	- I = 1
	- V = 5
	- X = 10
	- L = 50
	- C = 100
	- D = 500
	- M = 1000
*/

func main() {
	fmt.Println("\t-----ROMAN NUMERAL TO INTEGER-----")

	romanNumeral := "XXVII"
	convertedInt := intConversion(romanNumeral)

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Roman Numeral:\n", romanNumeral)
	fmt.Println("Integer Value:\n", convertedInt)
	fmt.Println(strings.Repeat("-", 50))
}

func intConversion(ri string) int {
	values := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(ri); i++ {
		currentVal := values[ri[i]]

		if i + 1 < len(ri) && currentVal < values[ri[i+1]] {
			total -= currentVal
		} else {
			total += currentVal
		}

	}

	return total
}
